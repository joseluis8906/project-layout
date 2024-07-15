package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"
	"github.com/joseluis8906/project-layout/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type (
	WkrDeps struct {
		fx.In
		Log         *log.Logger
		Kafka       *kafka.Conn
		RabbitMQ    *rabbitmq.Conn
		TxRepo      *Repository
		AccountRepo *account.Repository
	}

	Worker struct {
		LogPrintf      func(format string, v ...any)
		LogPrint       func(v ...any)
		KafkaPublish   func(topic string, msg []byte) error
		TxPersist      func(context.Context, Tx) error
		TxGet          func(ctx context.Context, txID string) (Tx, error)
		AccountGet     func(ctx context.Context, bank, atype, number string) (account.Account, error)
		AccountPersist func(context.Context, account.Account) error
	}
)

const (
	completedTransfersTopic = "banking.v1.completed_transfers"
	transfersQueue          = "banking.transfers"
)

func NewWorker(deps WkrDeps) *Worker {
	w := Worker{
		LogPrintf:      deps.Log.Printf,
		LogPrint:       deps.Log.Print,
		KafkaPublish:   deps.Kafka.Publish,
		TxGet:          deps.TxRepo.Get,
		TxPersist:      deps.TxRepo.Persist,
		AccountGet:     deps.AccountRepo.Get,
		AccountPersist: deps.AccountRepo.Persist,
	}

	deps.RabbitMQ.Subscribe(transfersQueue, w.ProcessTransfer)
	return &w
}

func (s *Worker) ProcessTransfer(d amqp.Delivery) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelFn()

	ctx, span := otel.Tracer("").Start(ctx, "banking.TxWorker/ProcessTransfer")
	defer span.End()

	var task pb.TransferJob
	err := proto.Unmarshal(d.Body, &task)
	if err != nil {
		s.LogPrintf("umarshaling message: %v", err)
		d.Reject(true)
		return
	}

	tx, err := s.TxGet(ctx, task.Id)
	if err != nil {
		s.LogPrintf("getting tx from repository: %v", err)
		d.Reject(true)
		return
	}

	var srcAccount, dstAccount account.Account
	g := new(errgroup.Group)
	g.Go(func() error {
		a, err := s.AccountGet(ctx, tx.SrcAccount.Bank, tx.SrcAccount.Type, tx.SrcAccount.Number)
		if err != nil {
			return fmt.Errorf("getting src account: %w", err)
		}
		srcAccount = a
		return nil
	})

	g.Go(func() error {
		a, err := s.AccountGet(ctx, tx.DstAccount.Bank, tx.DstAccount.Type, tx.DstAccount.Number)
		if err != nil {
			return fmt.Errorf("getting dst account: %w", err)
		}
		dstAccount = a
		return nil
	})

	if err = g.Wait(); err != nil {
		s.LogPrint(err)
		d.Reject(true)
		return
	}

	if !account.HasBalance(srcAccount, tx.Amount) {
		s.LogPrintf("source account does not have enough balance")
		d.Ack(false)
		return
	}

	if err := account.Debit(&srcAccount, tx.Amount); err != nil {
		s.LogPrintf("debitting amount: %v", err)
		d.Reject(true)
		return
	}

	if err := account.Credit(&dstAccount, tx.Amount); err != nil {
		s.LogPrintf("creditting account: %v", err)
		d.Reject(true)
		return
	}

	g.Go(func() error {
		if err := s.AccountPersist(ctx, srcAccount); err != nil {
			return fmt.Errorf("persisting src account: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		if err := s.AccountPersist(ctx, dstAccount); err != nil {
			return fmt.Errorf("persisting dst account: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		tx.Status = "completed"
		if err := s.TxPersist(ctx, tx); err != nil {
			return fmt.Errorf("persisting tx: %w", err)
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		s.LogPrint(err)
		d.Reject(true)
		return
	}

	if err := d.Ack(false); err != nil {
		s.LogPrintf("acknowledging message: %v", err)
		return
	}

	evt, err := proto.Marshal(&pb.Events_V1_TransferCompleted{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pb.Events_V1_TransferCompleted_Attributes{
			SrcAccount: &pb.Events_V1_TransferCompleted_Account{
				Bank:   srcAccount.Bank,
				Type:   srcAccount.Type,
				Number: srcAccount.Number,
			},
			DstAccount: &pb.Events_V1_TransferCompleted_Account{
				Bank:   dstAccount.Bank,
				Type:   dstAccount.Type,
				Number: dstAccount.Number,
			},
			Amount: &pkgpb.Money{
				Amount:   tx.Amount.Amount,
				Currency: tx.Amount.Currency,
			},
		},
	})
	if err != nil {
		s.LogPrintf("marshaling event: %v", err)
	}

	if err := s.KafkaPublish(completedTransfersTopic, evt); err != nil {
		s.LogPrintf("publishing event: %v", err)
	}
}
