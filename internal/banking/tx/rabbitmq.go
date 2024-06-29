package tx

import (
	"context"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type (
	WkrDeps struct {
		fx.In
		TxRepo      *Repository
		AccountRepo *account.Repository
	}

	Worker struct {
		Log         *log.Logger
		TxPersistor interface {
			Persist(context.Context, Tx) error
		}
		TxGetter interface {
			Get(context.Context, string) (Tx, error)
		}
		AccountGetter interface {
			Get(context.Context, string, string, string) (account.Account, error)
		}
		AccountPersistor interface {
			Persist(context.Context, account.Account) error
		}
	}
)

func NewWorker(deps WkrDeps) *Worker {
	return &Worker{
		TxGetter:         deps.TxRepo,
		TxPersistor:      deps.TxRepo,
		AccountGetter:    deps.AccountRepo,
		AccountPersistor: deps.AccountRepo,
	}
}

func (s *Worker) ProcessTransfer(d amqp.Delivery) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelFn()

	ctx, span := otel.Tracer("").Start(ctx, "banking.TxWorker/ProcessTransfer")
	defer span.End()

	var task pb.TransferJob
	err := proto.Unmarshal(d.Body, &task)
	if err != nil {
		s.Log.Printf("umarshaling message: %v", err)
		d.Reject(true)
		return
	}

	tx, err := s.TxGetter.Get(ctx, task.Id)
	if err != nil {
		s.Log.Printf("getting tx from repository: %v", err)
		d.Reject(true)
		return
	}

	var srcAccount, dstAccount account.Account
	g := new(errgroup.Group)
	g.Go(func() error {
		a, err := s.AccountGetter.Get(ctx, tx.SrcAccount.Bank, tx.SrcAccount.Type, tx.SrcAccount.Number)
		if err != nil {
			s.Log.Printf("getting src account: %v", err)
			d.Reject(true)
			return err
		}
		srcAccount = a
		return nil
	})

	g.Go(func() error {
		a, err := s.AccountGetter.Get(ctx, tx.DstAccount.Bank, tx.DstAccount.Type, tx.DstAccount.Number)
		if err != nil {
			s.Log.Printf("getting dst account: %v", err)
			d.Reject(true)
			return err
		}
		dstAccount = a
		return nil
	})

	if err = g.Wait(); err != nil {
		s.Log.Printf("getting src and dst accounts: %v", err)
		d.Reject(true)
		return
	}

	if !srcAccount.HasEnoughBalance(tx.Amount) {
		s.Log.Printf("source account does not have enough balance")
		d.Ack(false)
		return
	}

	if err := srcAccount.Debit(tx.Amount); err != nil {
		s.Log.Printf("debitting amount: %v", err)
		d.Reject(true)
		return
	}

	if err := dstAccount.Credit(tx.Amount); err != nil {
		s.Log.Printf("creditting account: %v", err)
		d.Reject(true)
		return
	}

	g.Go(func() error {
		if err := s.AccountPersistor.Persist(ctx, srcAccount); err != nil {
			s.Log.Printf("persisting src account: %v", err)
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := s.AccountPersistor.Persist(ctx, dstAccount); err != nil {
			s.Log.Printf("persisting dst account: %v", err)
			return err
		}
		return nil
	})

	g.Go(func() error {
		tx.Status = "completed"
		if err := s.TxPersistor.Persist(ctx, tx); err != nil {
			s.Log.Printf("persisting tx: %v", err)
			return err
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		s.Log.Printf("persisting src, dst account and tx")
		d.Reject(true)
		return
	}

	if err := d.Ack(false); err != nil {
		s.Log.Printf("acknowledging message: %v", err)
		return
	}
}
