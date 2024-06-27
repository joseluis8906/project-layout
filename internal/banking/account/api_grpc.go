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
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log     *log.Logger
		Kafka   *kafka.Conn
		Mongodb *mongo.Client
	}

	Service struct {
		pb.UnimplementedAccountServiceServer
		Log          *log.Logger
		Kafka        *kafka.Conn
		AccountAdder interface {
			Add(context.Context, Account) error
		}
		AccountGetter interface {
			Get(context.Context, string, string) (Account, error)
		}
	}
)

func New(deps Deps) *Service {
	accountRepo := NewAccountRepo(deps.Mongodb)
	s := &Service{
		Log:           deps.Log,
		Kafka:         deps.Kafka,
		AccountAdder:  accountRepo,
		AccountGetter: accountRepo,
	}

	return s
}

func (s *Service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	account := Account{
		Number:  fmt.Sprintf("%d", time.Now().Unix()),
		Type:    req.Type,
		Balance: money.New(0, req.Currency),
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

	err := s.AccountAdder.Add(ctx, account)
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
	account, err := s.AccountGetter.Get(ctx, req.Type, req.Number)
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

	if err = s.AccountAdder.Add(ctx, account); err != nil {
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
