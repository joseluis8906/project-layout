package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/money"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log      *log.Logger
		Kafka    *kafka.Conn
		RabbitMQ *amqp.Connection
		MongoDB  *mongo.Client
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

func New(deps Deps) *Service {
	txCh, err := deps.RabbitMQ.Channel()
	if err != nil {
		panic(fmt.Sprintf("creating amqp tx channel: %v", err))
	}
	s := &Service{
		Log:      deps.Log,
		Kafka:    deps.Kafka,
		RabbitMQ: txCh,
	}

	rxCh, err := deps.RabbitMQ.Channel()
	if err != nil {
		panic(fmt.Sprintf("creating amqp rx channel: %v", err))
	}

	msgs, err := rxCh.Consume("banking.init_txs", "", false, false, false, false, nil)
	if err != nil {
		panic(fmt.Sprintf("consuming amqp messages: %v", err))
	}

	go func() {
		w := Worker{
			Log: deps.Log,
		}

		for d := range msgs {
			w.ProcessInitTx(d)
		}
	}()

	return s
}

func (s *Service) InitTx(ctx context.Context, req *pb.InitTxRequest) (*pb.InitTxResponse, error) {
	tmCtx, cancelFn := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancelFn()

	txID := fmt.Sprintf("%d", time.Now().UnixMilli())
	tx := Tx{
		ID: txID,
		SrcBank: Bank{
			Name: req.SrcBank.Name,
			Account: account.Account{
				Type:   req.SrcBank.Account.Type,
				Number: req.SrcBank.Account.Number,
			},
		},
		DstBank: Bank{
			Name: req.DstBank.Name,
			Account: account.Account{
				Type:   req.DstBank.Account.Type,
				Number: req.DstBank.Account.Number,
			},
		},
		Amount: money.New(req.Amount.Amount, req.Amount.Currency),
		Status: "pending",
	}
	if err := s.TxPersistor.Persist(ctx, tx); err != nil {
		s.Log.Printf("persisting tx: %v", err)
		return nil, err
	}

	data, err := proto.Marshal(&pb.InitTxJob{
		Id:      txID,
		SrcBank: req.SrcBank,
		DstBank: req.DstBank,
		Amount:  req.Amount,
	})
	if err != nil {
		s.Log.Printf("marshaling task: %v", err)
		return nil, err
	}

	msg := amqp.Publishing{DeliveryMode: amqp.Persistent, Body: data}
	if err := s.RabbitMQ.PublishWithContext(tmCtx, "", "banking.init_txs", false, false, msg); err != nil {
		s.Log.Printf("publishing amqp message: %v", err)
		return nil, err
	}

	return &pb.InitTxResponse{TxId: txID}, nil
}

func (s *Service) CheckTxStatus(ctx context.Context, req *pb.CheckTxStatusRequest) (*pb.CheckTxStatusResponse, error) {
	tx, err := s.TxGetter.Get(ctx, req.TxId)
	if err != nil {
		s.Log.Printf("getting tx from repository: %v", err)
		return nil, err
	}

	return &pb.CheckTxStatusResponse{Status: tx.Status}, nil
}
