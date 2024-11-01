package tx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"go.uber.org/fx"

	"github.com/joseluis8906/project-layout/internal/bank/account"
	"github.com/joseluis8906/project-layout/internal/bank/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/money"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"
)

const (
	creditedAccountsTopic = "bank.v1.creditedAccounts"
	debitedAccountsTopic  = "bank.v1.debitedAccounts"
)

type (
	SvcDeps struct {
		fx.In
		Log         *log.Logger
		Kafka       *kafka.Conn
		AccountRepo *account.Repository
		TxRepo      *Repository
	}

	Service struct {
		pb.UnimplementedTxServiceServer
		log            *log.Logger
		kafka          *kafka.Conn
		AccountGet     func(ctx context.Context, kind, number string) (account.Account, error)
		AccountPersist func(context.Context, account.Account) error
		TxPersist      func(context.Context, Tx) error
	}
)

func NewGRPC(deps SvcDeps) *Service {
	s := &Service{
		log:            deps.Log,
		kafka:          deps.Kafka,
		AccountGet:     deps.AccountRepo.Get,
		AccountPersist: deps.AccountRepo.Persist,
		TxPersist:      deps.TxRepo.Persist,
	}

	return s
}

func (s *Service) Witdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	source, err := s.AccountGet(ctx, req.Account.Type, req.Account.Number)
	if err != nil {
		s.log.Printf("getting %s %s: %v", source.Type, source.Number, err)
		return nil, fmt.Errorf("getting %s %s: %w", req.Account.Type, req.Account.Number, err)
	}

	amount := money.New(req.Amount.Value, req.Amount.Currency)
	if err := account.Debit(&source, amount); err != nil {
		s.log.Printf("debiting %s %s: %v", source.Type, source.Number, err)
		return nil, fmt.Errorf("debiting amount for %s %s: %w", source.Type, source.Number, err)
	}

	if err := s.AccountPersist(ctx, source); err != nil {
		s.log.Printf("persisting %s %s: %v", source.Type, source.Number, err)
		return nil, fmt.Errorf("persisting %s %s: %w", source.Type, source.Number, err)
	}

	evt, err := proto.Marshal(&pb.Events_V1_AccountDebited{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pb.Events_V1_AccountDebited_Attributes{
			Type:   source.Type,
			Number: source.Number,
			Amount: &pkgpb.Money{Value: amount.Value, Currency: amount.Currency},
		},
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
		return nil, fmt.Errorf("marshaling event: %w", err)
	}

	if err := s.kafka.Publish(debitedAccountsTopic, evt); err != nil {
		s.log.Printf("publishing event: %v", err)
		return nil, fmt.Errorf("publishing event: %w", err)
	}

	return &pb.WithdrawResponse{Status: "success"}, nil
}

func (s *Service) DirectDeposit(ctx context.Context, req *pb.DirectDepositRequest) (*pb.DirectDepositResponse, error) {
	destintaion, err := s.AccountGet(ctx, req.Account.Type, req.Account.Number)
	if err != nil {
		return nil, fmt.Errorf("getting %s %s: %w", req.Account.Type, req.Account.Number, err)
	}

	amount := money.New(req.Amount.Value, req.Amount.Currency)
	if err := account.Credit(&destintaion, amount); err != nil {
		return nil, fmt.Errorf("crediting %s %s: %w", destintaion.Type, destintaion.Number, err)
	}

	if err := s.AccountPersist(ctx, destintaion); err != nil {
		s.log.Printf("persisting %s %s: %v", destintaion.Type, destintaion.Number, err)
		return nil, fmt.Errorf("persisting %s %s: %w", destintaion.Type, destintaion.Number, err)
	}

	evt, err := proto.Marshal(&pb.Events_V1_AccountCredited{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
		Attributes: &pb.Events_V1_AccountCredited_Attributes{
			Type:   destintaion.Type,
			Number: destintaion.Number,
			Amount: &pkgpb.Money{Value: amount.Value, Currency: amount.Currency},
		},
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
		return nil, fmt.Errorf("marshaling event: %w", err)
	}

	if err := s.kafka.Publish(creditedAccountsTopic, evt); err != nil {
		s.log.Printf("publishing event: %v", err)
		return nil, fmt.Errorf("publishing event: %w", err)
	}

	return &pb.DirectDepositResponse{Status: "success"}, nil
}
