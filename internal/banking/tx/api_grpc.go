package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	pkglog "github.com/joseluis8906/project-layout/pkg/log"
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
		Mongodb  *mongo.Client
	}

	Service struct {
		pb.UnimplementedTxServiceServer
		log      *log.Logger
		kafka    *kafka.Conn
		rabbitmq *amqp.Channel
		c        int
	}
)

func New(deps Deps) *Service {
	txCh, err := deps.RabbitMQ.Channel()
	if err != nil {
		panic(fmt.Sprintf("creating amqp tx channel: %v", err))
	}
	s := &Service{
		log:      deps.Log,
		kafka:    deps.Kafka,
		rabbitmq: txCh,
		c:        1,
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
		for d := range msgs {
			s.ProcessInitTx(d)
		}
	}()

	return s
}

func (s *Service) InitTx(ctx context.Context, req *pb.InitTxRequest) (*pb.InitTxResponse, error) {
	tmCtx, cancelFn := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancelFn()

	txID := fmt.Sprintf("%d", time.Now().UnixMilli())
	data, err := proto.Marshal(&pb.InitTxJob{
		Id:      txID,
		SrcBank: req.SrcBank,
		DstBank: req.DstBank,
		Amount:  req.Amount,
	})
	if err != nil {
		s.log.Printf("marshaling task: %v", err)
		return nil, err
	}

	msg := amqp.Publishing{DeliveryMode: amqp.Persistent, Body: data}
	if err := s.rabbitmq.PublishWithContext(tmCtx, "", "banking.init_txs", false, false, msg); err != nil {
		s.log.Printf("publishing amqp message: %v", err)
		return nil, err
	}

	return &pb.InitTxResponse{TxId: txID}, nil
}

func (s *Service) ProcessInitTx(d amqp.Delivery) {
	var task pb.InitTxJob
	err := proto.Unmarshal(d.Body, &task)
	if err != nil {
		s.log.Printf("umarshaling message: %v", err)
		return
	}

	s.log.Printf(pkglog.Info("task arrived: %s"), task.Id)
	time.Sleep(3 * time.Second)
	if s.c < 3 {
		s.log.Printf(pkglog.Info("task incompleted: %s"), task.Id)
		d.Reject(true)
		s.c++
		return
	}

	s.log.Printf(pkglog.Info("task completed: %s"), task.Id)
	if err := d.Ack(false); err != nil {
		s.log.Printf("acknowledging message: %v", err)
		return
	}
}
