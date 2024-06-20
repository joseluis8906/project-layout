package order

import (
	"context"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/pos/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	loglevel "github.com/joseluis8906/project-layout/pkg/log"
	evtpb "github.com/joseluis8906/project-layout/pkg/pb"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log   *log.Logger
		Kafka *kafka.Conn
	}

	Service struct {
		pb.UnimplementedOrderServiceServer
		log        *log.Logger
		kafka      *kafka.Conn
		OrderAdder interface {
			Add(context.Context, Order) error
		}
		Counter interface {
			Count() uint64
		}
	}
)

func New(deps Deps) *Service {
	s := &Service{
		log:   deps.Log,
		kafka: deps.Kafka,
	}

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
	return s
}

func (s *Service) Create(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	evt, err := proto.Marshal(&evtpb.V1_Tested{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Msg:        req.Customer,
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
	}

	err = s.kafka.Publish("v1.tested", evt)
	if err != nil {
		s.log.Printf("publishing event: %v", err)
	}

	return &pb.CreateOrderResponse{Number: 1}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Tracer("").Start(context.Background(), "pos.OrderService/OnTested")
	defer span.End()

	var evt evtpb.V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.log.Printf("unmarshaling event: %v", err)
		return
	}

	s.log.Printf(loglevel.Info(`msg received: {"id": %s, "occurred_on": %s, "msg": %s}`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Msg)
}
