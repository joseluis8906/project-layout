package account

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/joseluis8906/project-layout/pkg/money"
)

type (
	Account struct {
		PhoneNumber string      `bson:"phone_number"`
		Balance     money.Money `bson:"balance"`
		Owner       Owner       `bson:"owner"`
	}

	Owner struct {
		ID       string `bson:"id"`
		Email    string `bson:"email"`
		FullName string `bson:"full_name"`
	}
)

func Validate(a Account) error {
	validPhone := regexp.MustCompile(`^\+(57|52|593)\d{10}$`)
	validID := regexp.MustCompile(`^\d+$`)
	validEmail := regexp.MustCompile(`^[\w]+.*@[\w]+.(com|net|org)`)
	validName := regexp.MustCompile(`^[\w]{2,} [\w]{2,}$`)
	if !validPhone.MatchString(a.PhoneNumber) {
		return fmt.Errorf("invalid phone number %q", a.PhoneNumber)
	}

	if !validID.MatchString(a.Owner.ID) {
		return errors.New("invalid owner id")
	}

	if !validEmail.MatchString(a.Owner.Email) {
		return errors.New("invalid owner email")
	}

	if !validName.MatchString(a.Owner.FullName) {
		return errors.New("invalid owner name")
	}

	return nil
}

func PutMoney(account Account, amount money.Money) {
	account.Balance = money.Add(account.Balance, amount)
}

func SendMoney(srcAccount, dstAccount *Account, amount money.Money) error {
	if err := Debit(srcAccount, amount); err != nil {
		return fmt.Errorf("debiting amount: %w", err)
	}

	Credit(dstAccount, amount)
	return nil
}

func Debit(account *Account, amount money.Money) error {
	if !hasBalance(*account, amount) {
		return errors.New("account without balance")
	}

	return nil
}

func Credit(account *Account, amount money.Money) {
	account.Balance = money.Add(account.Balance, amount)
}

func hasBalance(account Account, amount money.Money) bool {
	return account.Balance.Value >= amount.Value
}
