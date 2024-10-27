package account_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/joseluis8906/project-layout/internal/bank/account"
	"github.com/joseluis8906/project-layout/internal/bank/pb"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"
)

func TestService_CreateAccount(t *testing.T) {
	svc := account.Service{
		LogPrintf:      func(format string, v ...any) {},
		KafkaPublish:   func(topic string, msg []byte) error { return nil },
		AccountPersist: func(ctx context.Context, a account.Account) error { return nil },
		AccountGet: func(ctx context.Context, atype, number string) (account.Account, error) {
			return account.Account{}, nil
		},
	}

	input := &pb.CreateAccountRequest{
		Type: account.TypeSaving,
		Owner: &pb.CreateAccountRequest_Owner{
			Id:       "10",
			Email:    "john.doe@example.com",
			Country:  account.CountryCO,
			FullName: "John Doe",
		},
		Balance: &pkgpb.Money{
			Value:    1000,
			Currency: "USD",
		},
	}
	got, err := svc.CreateAccount(context.TODO(), input)
	want := &pb.CreateAccountResponse{Number: got.Number}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(pb.CreateAccountResponse{})); err != nil || diff != "" {
		t.Errorf("account.Service.CreateAccount(ctx, %v) = %v, %v; want %v, <nil>\n(-want, +got)\n%s", input, got, err, want, diff)
	}
}
