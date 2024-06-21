package account

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joseluis8906/project-layout/internal/banking/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	evtpb "github.com/joseluis8906/project-layout/pkg/pb"
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
		log       *log.Logger
		kafka     *kafka.Conn
		accounter interface {
			Add(context.Context, Account) error
		}
	}
)

func New(deps Deps) *Service {
	s := &Service{
		log:       deps.Log,
		kafka:     deps.Kafka,
		accounter: NewAccountRepo(deps.Mongodb),
	}

	return s
}

func (s *Service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	account := Account{
		Type:   "saving account",
		Number: fmt.Sprintf("%d", time.Now().Unix()),
		Owner: Owner{
			ID:       req.Id,
			Email:    req.Email,
			Country:  req.Country,
			FullName: req.FullName,
		},
	}

	if err := account.Validate(); err != nil {
		log.Printf("validating account: %v", err)
		return nil, fmt.Errorf("validating account owner: %w", err)
	}

	err := s.accounter.Add(ctx, account)
	if err != nil {
		log.Printf("adding account: %v", err)
		return nil, fmt.Errorf("adding account: %w", err)
	}

	evt, err := proto.Marshal(&evtpb.V1_AccountCreated{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &evtpb.V1_AccountCreated_Attributes{
			Type:   account.Type,
			Number: account.Number,
		},
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
	}

	err = s.kafka.Publish("v1.account_created", evt)
	if err != nil {
		s.log.Printf("publishing event: %v", err)
	}

	return &pb.CreateAccountResponse{Type: account.Type, Number: account.Number}, nil
}
