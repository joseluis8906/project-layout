package account

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/bank/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/metric"
	"github.com/joseluis8906/project-layout/pkg/money"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	SvcDeps struct {
		fx.In
		Log         *log.Logger
		Metric      *metric.Collector
		Kafka       *kafka.Conn
		AccountRepo *Repository
	}

	Service struct {
		pb.UnimplementedAccountServiceServer
		log     *log.Logger
		metric  *metric.Collector
		kafka   *kafka.Conn
		Account struct {
			Persist func(context.Context, Account) error
		}
	}
)

const (
	createdAccountsTopic = "bank.v1.created_accounts"
)

func NewGRPC(deps SvcDeps) *Service {
	s := &Service{
		log:    deps.Log,
		metric: deps.Metric,
		kafka:  deps.Kafka,
	}

	s.Account.Persist = deps.AccountRepo.Persist

	return s
}

func (s *Service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (res *pb.CreateAccountResponse, err error) {
	defer func() {
		s.metric.OpsResult(err, "bank.AccountService", "CreateAccount")
	}()

	newAccount := Account{
		Type:    req.Type,
		Number:  fmt.Sprintf("%d", time.Now().Unix()),
		Balance: money.New(req.Balance.Value, req.Balance.Currency),
		Owner: Owner{
			ID:       req.Owner.Id,
			Email:    req.Owner.Email,
			Country:  req.Owner.Country,
			FullName: req.Owner.FullName,
		},
	}
	if err = Validate(newAccount); err != nil {
		s.log.Printf("validating account: %v", err)
		return nil, fmt.Errorf("validating account owner: %w", err)
	}

	err = s.Account.Persist(ctx, newAccount)
	if err != nil {
		s.log.Printf("adding account: %v", err)
		return nil, fmt.Errorf("adding account: %w", err)
	}

	evt, err := proto.Marshal(&pb.Events_V1_AccountCreated{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pb.Events_V1_AccountCreated_Attributes{
			Type:   newAccount.Type,
			Number: newAccount.Number,
		},
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
		return nil, fmt.Errorf("marshaling event: %w", err)
	}

	if err := s.kafka.Publish(createdAccountsTopic, evt); err != nil {
		s.log.Printf("publishing event: %v", err)
		return nil, fmt.Errorf("publishing event: %w", err)
	}

	return &pb.CreateAccountResponse{Number: newAccount.Number}, nil
}
