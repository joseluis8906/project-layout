package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/money"
	"github.com/joseluis8906/project-layout/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	SvcDeps struct {
		fx.In
		Log      *log.Logger
		Kafka    *kafka.Conn
		RabbitMQ *rabbitmq.Conn
		TxRepo   *Repository
		Worker   *Worker
	}

	Service struct {
		pb.UnimplementedTxServiceServer
		LogPrintf       func(format string, v ...any)
		Kafka           *kafka.Conn
		RabbitMQPublish func(ctx context.Context, exchange, key string, mandatory, inmediate bool, msg amqp.Publishing) error
		TxPersist       func(context.Context, Tx) error
		TxGet           func(context.Context, string) (Tx, error)
	}
)

const (
	transfersQueue = "banking.transfers"
)

func New(deps SvcDeps) *Service {
	s := &Service{
		LogPrintf:       deps.Log.Printf,
		Kafka:           deps.Kafka,
		RabbitMQPublish: deps.RabbitMQ.Publish,
		TxPersist:       deps.TxRepo.Persist,
		TxGet:           deps.TxRepo.Get,
	}

	deps.RabbitMQ.Subscribe(transfersQueue, deps.Worker.ProcessTransfer)
	return s
}

func (s *Service) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	tmCtx, cancelFn := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancelFn()

	txID := fmt.Sprintf("%d", time.Now().UnixMilli())
	newTx := Tx{
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
	if err := Validate(newTx); err != nil {
		s.LogPrintf("validiting tx: %v", err)
		return nil, err
	}

	if err := s.TxPersist(ctx, newTx); err != nil {
		s.LogPrintf("persisting tx: %v", err)
		return nil, err
	}

	data, err := proto.Marshal(&pb.TransferJob{
		Id:         txID,
		Amount:     req.Amount,
		SrcAccount: req.SrcAccount,
		DstAccount: req.DstAccount,
	})
	if err != nil {
		s.LogPrintf("marshaling task: %v", err)
		return nil, err
	}

	msg := amqp.Publishing{DeliveryMode: amqp.Persistent, Body: data}
	//
	if err := s.RabbitMQPublish(tmCtx, "", transfersQueue, false, false, msg); err != nil {
		s.LogPrintf("publishing amqp message: %v", err)
		return nil, err
	}

	return &pb.TransferResponse{TxId: txID}, nil
}

func (s *Service) CheckTxStatus(ctx context.Context, req *pb.CheckTxStatusRequest) (*pb.CheckTxStatusResponse, error) {
	tx, err := s.TxGet(ctx, req.TxId)
	if err != nil {
		s.LogPrintf("getting tx from repository: %v", err)
		return nil, err
	}

	return &pb.CheckTxStatusResponse{Status: tx.Status}, nil
}
