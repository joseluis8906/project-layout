package account

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/mtx/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	pkglog "github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/money"
	"github.com/joseluis8906/project-layout/pkg/otel"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"

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
		LogPrintf      func(format string, v ...any)
		AccountPersist func(context.Context, Account) error
	}
)

func New(deps Deps) *Service {
	s := &Service{
		LogPrintf:      deps.Log.Printf,
		AccountPersist: deps.AccountRepo.Persist,
	}

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
	return s
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newAccount := Account{
		PhoneNumber: req.PhoneNumber,
		Balance:     money.New(0, money.COP),
		Owner: Owner{
			ID:       req.Owner.Id,
			Email:    req.Owner.Email,
			FullName: req.Owner.FullName,
		},
	}
	if err := Validate(newAccount); err != nil {
		s.LogPrintf("validating account: %v", err)
		return nil, fmt.Errorf("validating account: %w", err)
	}

	if err := s.AccountPersist(ctx, newAccount); err != nil {
		s.LogPrintf("persisting account: %v", err)
		return nil, fmt.Errorf("persisting account: %w", err)
	}

	return &pb.RegisterResponse{}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Start(context.Background(), otel.NoTracer, "mtx.HelloService/OnTested")
	defer span.End()

	var evt pkgpb.V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.LogPrintf("unmarshaling event: %v", err)
		return
	}

	s.LogPrintf(pkglog.Info(`msg received: {"id": %s, "occurred_on": %s, "msg": %s}`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Msg)
}
