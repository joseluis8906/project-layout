package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/money"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	SvcDeps struct {
		fx.In
		Log      *log.Logger
		Kafka    *kafka.Conn
		RabbitMQ *amqp.Connection
		TxRepo   *Repository
		Worker   *Worker
	}

	Service struct {
		pb.UnimplementedTxServiceServer
		Log         *log.Logger
		Kafka       *kafka.Conn
		RabbitMQ    *amqp.Channel
		TxPersistor interface {
			Persist(context.Context, Tx) error
		}
		TxGetter interface {
			Get(context.Context, string) (Tx, error)
		}
	}
)

func New(deps SvcDeps) *Service {
	txCh, err := deps.RabbitMQ.Channel()
	if err != nil {
		panic(fmt.Sprintf("creating amqp tx channel: %v", err))
	}
	s := &Service{
		Log:         deps.Log,
		Kafka:       deps.Kafka,
		RabbitMQ:    txCh,
		TxPersistor: deps.TxRepo,
		TxGetter:    deps.TxRepo,
	}

	rxCh, err := deps.RabbitMQ.Channel()
	if err != nil {
		panic(fmt.Sprintf("creating amqp rx channel: %v", err))
	}

	msgs, err := rxCh.Consume("banking.transfers", "", false, false, false, false, nil)
	if err != nil {
		panic(fmt.Sprintf("consuming amqp messages: %v", err))
	}

	go func() {
		for d := range msgs {
			deps.Worker.ProcessTransfer(d)
		}
	}()

	return s
}

func (s *Service) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	tmCtx, cancelFn := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancelFn()

	txID := fmt.Sprintf("%d", time.Now().UnixMilli())
	tx := Tx{
		ID:     txID,
		Amount: money.New(req.Amount.Amount, req.Amount.Currency),
		SrcAccount: Account{
			Bank:   req.SrcAccount.Bank,
			Type:   req.SrcAccount.Type,
			Number: req.SrcAccount.Number,
		},
		DstAccount: Account{
			Bank:   req.DstAccount.Bank,
			Type:   req.DstAccount.Type,
			Number: req.DstAccount.Number,
		},
		Status: "pending",
	}
	if err := tx.Validate(); err != nil {
		s.Log.Printf("validiting tx: %v", err)
		return nil, err
	}

	if err := s.TxPersistor.Persist(ctx, tx); err != nil {
		s.Log.Printf("persisting tx: %v", err)
		return nil, err
	}

	data, err := proto.Marshal(&pb.TransferJob{
		Id:         txID,
		Amount:     req.Amount,
		SrcAccount: req.SrcAccount,
		DstAccount: req.DstAccount,
	})
	if err != nil {
		s.Log.Printf("marshaling task: %v", err)
		return nil, err
	}

	msg := amqp.Publishing{DeliveryMode: amqp.Persistent, Body: data}
	if err := s.RabbitMQ.PublishWithContext(tmCtx, "", "banking.transfers", false, false, msg); err != nil {
		s.Log.Printf("publishing amqp message: %v", err)
		return nil, err
	}

	return &pb.TransferResponse{TxId: txID}, nil
}

func (s *Service) CheckTxStatus(ctx context.Context, req *pb.CheckTxStatusRequest) (*pb.CheckTxStatusResponse, error) {
	tx, err := s.TxGetter.Get(ctx, req.TxId)
	if err != nil {
		s.Log.Printf("getting tx from repository: %v", err)
		return nil, err
	}

	return &pb.CheckTxStatusResponse{Status: tx.Status}, nil
}
