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
			Persist func(context.Context, *Account) error
			Get     func(ctx context.Context, phoneNumber string) (Account, error)
		}
	}
)

func NewGRPC(deps Deps) *Service {
	s := &Service{
		log:    deps.Log,
		metric: deps.Metric,
	}
	s.Account.Persist = deps.AccountRepo.Persist
	s.Account.Get = deps.AccountRepo.Get

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
	return s
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (res *pb.RegisterResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "mtx.AccountService", "Register")
	}()

	s.log.Printf("register")

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

	if err := s.Account.Persist(ctx, &newAccount); err != nil {
		s.log.Printf("persisting account: %v", err)
		return nil, fmt.Errorf("persisting account: %w", err)
	}

	return &pb.RegisterResponse{}, nil
}

func (s *Service) PutMoney(ctx context.Context, req *pb.PutMoneyRequest) (res *pb.PutMoneyResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "mtx.AccountService", "PutMoney")
	}()

	s.log.Printf("put money")
	account, err := s.Account.Get(ctx, req.GetPhoneNumber())
	if err != nil {
		return nil, fmt.Errorf("getting account: %w", err)
	}

	amount := money.New(req.GetAmount().GetValue(), req.GetAmount().GetCurrency())
	PutMoney(account, amount)

	if err := s.Account.Persist(ctx, &account); err != nil {
		return nil, fmt.Errorf("persisting account: %w", err)
	}

	return &pb.PutMoneyResponse{}, nil
}

func (s *Service) SendMoney(ctx context.Context, req *pb.SendMoneyRequest) (res *pb.SendMoneyResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "mtx.AccountService", "SendMoney")
	}()

	s.log.Printf("send money")
	srcAccount, err := s.Account.Get(ctx, req.GetSrcPhoneNumber())
	if err != nil {
		return nil, fmt.Errorf("getting source account: %w", err)
	}

	dstAccount, err := s.Account.Get(ctx, req.GetDstPhoneNumber())
	if err != nil {
		return nil, fmt.Errorf("getting dest account: %w", err)
	}

	amount := money.New(req.GetAmount().GetValue(), req.GetAmount().GetCurrency())
	if err := SendMoney(&srcAccount, &dstAccount, amount); err != nil {
		return nil, fmt.Errorf("sending money: %w", err)
	}

	if err := s.Account.Persist(ctx, &srcAccount); err != nil {
		return nil, fmt.Errorf("persisting source account: %w", err)
	}

	if err := s.Account.Persist(ctx, &dstAccount); err != nil {
		return nil, fmt.Errorf("persisting dest account: %w", err)
	}

	return &pb.SendMoneyResponse{}, nil
}

func (s *Service) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (res *pb.WithdrawResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "mtx.AccountService", "Withdraw")
	}()

	s.log.Printf("withdraw")
	account, err := s.Account.Get(ctx, req.GetPhoneNumber())
	if err != nil {
		return nil, fmt.Errorf("getting account: %w", err)
	}

	amount := money.New(req.GetAmount().GetValue(), req.GetAmount().GetCurrency())
	if err := Debit(&account, amount); err != nil {
		return nil, fmt.Errorf("debiting account: %w", err)
	}

	if err := s.Account.Persist(ctx, &account); err != nil {
		return nil, fmt.Errorf("persisting account: %w", err)
	}

	return &pb.WithdrawResponse{}, nil
}

func (s *Service) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (res *pb.GetBalanceResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "mtx.AccountService", "GetBalance")
	}()

	s.log.Printf("get balance")
	account, err := s.Account.Get(ctx, req.GetPhoneNumber())
	if err != nil {
		return nil, fmt.Errorf("getting account: %w", err)
	}

	return &pb.GetBalanceResponse{Balance: &pkgpb.Money{Value: account.Balance.Value, Currency: account.Balance.Currency}}, nil
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
