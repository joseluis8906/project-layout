package account_test

import (
	"errors"
	"testing"

	"github.com/joseluis8906/project-layout/internal/bank/account"
)

func TestValidate(t *testing.T) {
	testCases := map[string]struct {
		account account.Account
		want    error
	}{
		"valid": {
			account: account.Account{
				Type: "saving account",
				Owner: account.Owner{
					ID:       "123456789",
					FullName: "John Doe",
					Email:    "john.doe@example.com",
					Country:  "CO",
				},
			},
			want: nil,
		},
		"invalid id": {
			account: account.Account{
				Type: "saving account",
				Owner: account.Owner{
					ID: "a123456789",
				},
			},
			want: errors.New("invalid owner id"),
		},
		"invalid email": {
			account: account.Account{
				Type: "saving account",
				Owner: account.Owner{
					ID:    "123456789",
					Email: "jhon.doe@example",
				},
			},
			want: errors.New("invalid owner email"),
		},
		"invalid name": {
			account: account.Account{
				Type: "saving account",
				Owner: account.Owner{
					ID:       "123456789",
					Email:    "jhon.doe@example.com",
					FullName: "J Doe",
				},
			},
			want: errors.New("invalid owner name"),
		},
		"invalid country": {
			account: account.Account{
				Type: "saving account",
				Owner: account.Owner{
					ID:       "123456789",
					FullName: "John Doe",
					Email:    "jhon.doe@example.com",
					Country:  "US",
				},
			},
			want: errors.New("invalid owner's country"),
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			err := account.Validate(tc.account)
			if err != nil && tc.want.Error() != err.Error() {
				t.Errorf("account.Validate(%v) = %v, want %v", tc.account, err, tc.want)
			}
		})
	}
}
