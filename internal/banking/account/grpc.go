package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/money"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	SvcDeps struct {
		fx.In
		Log         *log.Logger
		Kafka       *kafka.Conn
		AccountRepo *Repository
	}

	Service struct {
		pb.UnimplementedAccountServiceServer
		Log              *log.Logger
		Kafka            *kafka.Conn
		AccountPersistor interface {
			Persist(context.Context, Account) error
		}
		AccountGetter interface {
			Get(context.Context, string, string, string) (Account, error)
		}
	}
)

func New(deps SvcDeps) *Service {
	s := &Service{
		Log:              deps.Log,
		Kafka:            deps.Kafka,
		AccountPersistor: deps.AccountRepo,
		AccountGetter:    deps.AccountRepo,
	}

	return s
}

func (s *Service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	account := Account{
		Bank:    req.Bank,
		Type:    req.Type,
		Number:  fmt.Sprintf("%d", time.Now().Unix()),
		Balance: money.New(req.Balance.Amount, req.Balance.Currency),
		Owner: Owner{
			ID:       req.Owner.Id,
			Email:    req.Owner.Email,
			Country:  req.Owner.Country,
			FullName: req.Owner.FullName,
		},
	}
	if err := account.Validate(); err != nil {
		log.Printf("validating account: %v", err)
		return nil, fmt.Errorf("validating account owner: %w", err)
	}

	err := s.AccountPersistor.Persist(ctx, account)
	if err != nil {
		log.Printf("adding account: %v", err)
		return nil, fmt.Errorf("adding account: %w", err)
	}

	evt, err := proto.Marshal(&pkgpb.V1_AccountCreated{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pkgpb.V1_AccountCreated_Attributes{
			Type:   account.Type,
			Number: account.Number,
		},
	})
	if err != nil {
		s.Log.Printf("marshaling event: %v", err)
	}

	err = s.Kafka.Publish("v1.account_created", evt)
	if err != nil {
		s.Log.Printf("publishing event: %v", err)
	}

	return &pb.CreateAccountResponse{Number: account.Number}, nil
}

func (s *Service) CreditAccount(ctx context.Context, req *pb.CreditAccountRequest) (*pb.CreditAccountResponse, error) {
	account, err := s.AccountGetter.Get(ctx, req.Bank, req.Type, req.Number)
	if err != nil {
		return nil, fmt.Errorf("getting account: %w", err)
	}

	if account.IsZero() {
		return nil, errors.New("account does not exist")
	}

	amount := money.New(req.Amount.Amount, req.Amount.Currency)
	if err := account.Credit(amount); err != nil {
		return nil, fmt.Errorf("crediting account: %w", err)
	}

	if err = s.AccountPersistor.Persist(ctx, account); err != nil {
		return nil, fmt.Errorf("updating account: %w", err)
	}

	evt, err := proto.Marshal(&pkgpb.V1_AccountCredited{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pkgpb.V1_AccountCredited_Attributes{
			Type:   account.Type,
			Number: account.Number,
			Amount: fmt.Sprintf("%s", amount),
		},
	})
	if err != nil {
		s.Log.Printf("marshaling event: %v", err)
	}

	if err := s.Kafka.Publish("v1.account_credited", evt); err != nil {
		s.Log.Printf("publishing event: %v", err)
	}

	return &pb.CreditAccountResponse{}, nil
}
