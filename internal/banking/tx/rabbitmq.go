package tx

import (
	"context"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/internal/banking/pb"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type (
	Worker struct {
		Log         *log.Logger
		TxPersistor interface {
			Persist(context.Context, Tx) error
		}
		TxGetter interface {
			Get(context.Context, string) (Tx, error)
		}
		AccountGetter interface {
			Get(context.Context, string, string) (account.Account, error)
		}
		AccountPersistor interface {
			Persist(context.Context, account.Account) error
		}
	}
)

func (s *Worker) ProcessInitTx(d amqp.Delivery) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancelFn()

	var task pb.InitTxJob
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

	srcAccount, err := s.AccountGetter.Get(ctx, tx.SrcBank.Account.Type, tx.SrcBank.Account.Number)
	if err != nil {
		s.Log.Printf("getting src account: %v", err)
		d.Reject(true)
		return
	}

	if !srcAccount.HasEnoughBalance(tx.Amount) {
		s.Log.Printf("source account does not have enough balance")
		d.Ack(false)
		return
	}

	dstAccount, err := s.AccountGetter.Get(ctx, tx.DstBank.Account.Type, tx.DstBank.Account.Number)
	if err != nil {
		s.Log.Printf("getting dst account: %v", err)
		d.Reject(true)
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

	if err := s.AccountPersistor.Persist(ctx, srcAccount); err != nil {
		s.Log.Printf("persisting src account: %v", err)
		return
	}

	if err := s.AccountPersistor.Persist(ctx, dstAccount); err != nil {
		s.Log.Printf("persisting dst account: %v", err)
		return
	}

	tx.Status = "completed"
	if err := s.TxPersistor.Persist(ctx, tx); err != nil {
		s.Log.Printf("persisting tx: %v", err)
		return
	}

	if err := d.Ack(false); err != nil {
		s.Log.Printf("acknowledging message: %v", err)
		return
	}
}
