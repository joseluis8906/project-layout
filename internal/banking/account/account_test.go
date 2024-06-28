package account_test

import (
	"errors"
	"testing"

	"github.com/joseluis8906/project-layout/internal/banking/account"
)

func TestAccount_Validate(t *testing.T) {
	testCases := map[string]struct {
		account account.Account
		want    error
	}{
		"valid": {
			account: account.Account{
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
				Owner: account.Owner{
					ID: "a123456789",
				},
			},
			want: errors.New("invalid id"),
		},
		"invalid email": {
			account: account.Account{
				Owner: account.Owner{
					ID:    "123456789",
					Email: "jhon.doe@example",
				},
			},
			want: errors.New("invalid email"),
		},
		"invalid name": {
			account: account.Account{
				Owner: account.Owner{
					ID:       "123456789",
					Email:    "jhon.doe@example.com",
					FullName: "J Doe",
				},
			},
			want: errors.New("invalid name"),
		},
		"invalid country": {
			account: account.Account{
				Owner: account.Owner{
					ID:       "123456789",
					FullName: "John Doe",
					Email:    "jhon.doe@example.com",
					Country:  "US",
				},
			},
			want: errors.New("invalid country"),
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			err := tc.account.Validate()
			if err != nil && tc.want.Error() != err.Error() {
				t.Errorf("account.Account.Validate() = %v, want %v", err, tc.want)
			}
		})
	}
}
