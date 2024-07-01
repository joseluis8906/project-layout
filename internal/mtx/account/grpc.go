package account

import (
	"context"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/mtx/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	pkglog "github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/money"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"

	// "github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log         *log.Logger
		Kafka       *kafka.Conn
		AccountRepo *Repository
	}

	Service struct {
		pb.UnimplementedAccountServiceServer
		log              *log.Logger
		kafka            *kafka.Conn
		AccountPersistor interface {
			Persist(context.Context, Account) error
		}
	}
)

func New(deps Deps) *Service {
	s := &Service{
		log:              deps.Log,
		kafka:            deps.Kafka,
		AccountPersistor: deps.AccountRepo,
	}

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
	return s
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	a := Account{
		PhoneNumber: req.PhoneNumber,
		Balance:     money.New(0, money.COP),
		Owner: Owner{
			ID:       req.Owner.Id,
			Email:    req.Owner.Email,
			FullName: req.Owner.FullName,
		},
	}
	if err := a.Validate(); err != nil {
		s.log.Printf("validating account: %v", err)
		return nil, err
	}

	if err := s.AccountPersistor.Persist(ctx, a); err != nil {
		s.log.Printf("persisting account: %v", err)
		return nil, err
	}

	return &pb.RegisterResponse{}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Tracer("").Start(context.Background(), "mtx.HelloService/OnTested")
	defer span.End()

	var evt pkgpb.V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.log.Printf("unmarshaling event: %v", err)
		return
	}

	s.log.Printf(pkglog.Info(`msg received: {"id": %s, "occurred_on": %s, "msg": %s}`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Msg)
}
