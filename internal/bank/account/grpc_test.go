package account_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/joseluis8906/project-layout/internal/bank/account"
	"github.com/joseluis8906/project-layout/internal/bank/pb"
	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/metric"
	pkgpb "github.com/joseluis8906/project-layout/pkg/pb"
)

func TestService_CreateAccount(t *testing.T) {
	svc := account.NewGRPC(account.SvcDeps{
		Log:    log.Noop(),
		Metric: metric.Noop(),
		Kafka:  kafka.Noop(),
	})

	svc.Account.Persist = func(ctx context.Context, a account.Account) error { return nil }

	in := &pb.CreateAccountRequest{
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

	got, err := svc.CreateAccount(context.TODO(), in)
	want := &pb.CreateAccountResponse{Number: got.Number}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(pb.CreateAccountResponse{})); err != nil || diff != "" {
		t.Errorf("account.Service.CreateAccount(ctx, %v) = %v, %v; want %v, <nil>\n(-want, +got)\n%s", in, got, err, want, diff)
	}
}
