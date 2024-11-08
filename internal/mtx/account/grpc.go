package account

import (
	"context"
	"fmt"
	stdlog "log"
	"time"

	"github.com/joseluis8906/project-layout/internal/mtx/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/metric"
	"github.com/joseluis8906/project-layout/pkg/money"
	"github.com/joseluis8906/project-layout/pkg/otel"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"

	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log         *stdlog.Logger
		Kafka       *kafka.Conn
		Metric      *metric.Collector
		AccountRepo *Repository
	}

	Service struct {
		pb.UnimplementedAccountServiceServer
		metric  *metric.Collector
		log     *stdlog.Logger
		Account struct {
			Persist func(context.Context, Account) error
		}
	}
)

func NewGRPC(deps Deps) *Service {
	s := &Service{
		log:    deps.Log,
		metric: deps.Metric,
	}
	s.Account.Persist = deps.AccountRepo.Persist

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
	return s
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterResponse, err error) {
	defer func() {
		s.metric.OpsResult(
			err,
			metric.Tag(metric.ServiceTagKey, "mtx.AccountService"),
			metric.Tag(metric.MethodTagKey, "Register"),
		)
	}()

	newAccount := Account{
		PhoneNumber: req.PhoneNumber,
		Balance:     money.New(0, money.COP),
		Owner: Owner{
			ID:       req.Owner.Id,
			Email:    req.Owner.Email,
			FullName: req.Owner.FullName,
		},
	}
	if err = Validate(newAccount); err != nil {
		s.log.Printf("validating account: %v", err)
		return nil, fmt.Errorf("validating account: %w", err)
	}

	if err := s.Account.Persist(ctx, newAccount); err != nil {
		s.log.Printf("persisting account: %v", err)
		return nil, fmt.Errorf("persisting account: %w", err)
	}

	return &pb.RegisterResponse{}, nil
}

func (s *Service) PutMoney(ctx context.Context, req *pb.PutMoneyRequest) (res *pb.PutMoneyResponse, err error) {
	defer func() {
		s.metric.OpsResult(
			err,
			metric.Tag(metric.ServiceTagKey, "mtx.AccountService"),
			metric.Tag(metric.MethodTagKey, "PutMoney"),
		)
	}()

	return &pb.PutMoneyResponse{Id: "testing"}, nil
}

func (s *Service) SendMoney(ctx context.Context, req *pb.SendMoneyRequest) (res *pb.SendMoneyResponse, err error) {
	defer func() {
		s.metric.OpsResult(
			err,
			metric.Tag(metric.ServiceTagKey, "mtx.AccountService"),
			metric.Tag(metric.MethodTagKey, "SendMoney"),
		)
	}()

	return &pb.SendMoneyResponse{Status: "testing"}, nil
}

func (s *Service) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (res *pb.WithdrawResponse, err error) {
	defer func() {
		s.metric.OpsResult(
			err,
			metric.Tag(metric.ServiceTagKey, "mtx.AccountService"),
			metric.Tag(metric.MethodTagKey, "Withdraw"),
		)
	}()

	return &pb.WithdrawResponse{Status: "testing"}, nil
}

func (s *Service) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (res *pb.GetBalanceResponse, err error) {
	defer func() {
		s.metric.OpsResult(
			err,
			metric.Tag(metric.ServiceTagKey, "mtx.AccountService"),
			metric.Tag(metric.MethodTagKey, "GetBalance"),
		)
	}()

	return &pb.GetBalanceResponse{Balance: &pkgpb.Money{}}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Start(context.Background(), otel.NoTracer, "mtx.HelloService/OnTested")
	defer span.End()

	var evt pkgpb.V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.log.Printf("unmarshaling event: %v", err)
		return
	}

	s.log.Printf(log.Info(`msg received: {"id": %s, "occurred_on": %s, "msg": %s}`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Msg)
}
