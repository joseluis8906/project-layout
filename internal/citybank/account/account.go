package account

import (
	"errors"
	"regexp"

	"github.com/joseluis8906/project-layout/pkg/money"
)

type (
	Account struct {
		Type    string      `bson:"type"`
		Number  string      `bson:"number"`
		Balance money.Money `bson:"balance"`
		Owner   Owner       `bson:"owner"`
	}

	Owner struct {
		ID       string `bson:"id"`
		Country  string `bson:"country"`
		Email    string `bson:"email"`
		FullName string `bson:"full_name"`
	}
)

func (a Account) IsZero() bool {
	return len(a.Type) == 0 || len(a.Number) == 0
}

func Validate(a Account) error {
	validID := regexp.MustCompile(`^[0-9]+$`)
	validEmail := regexp.MustCompile(`^[\w]+.*@[\w]+.(com|net|org)`)
	validName := regexp.MustCompile(`^[\w]{2,} [\w]{2,}$`)
	validCountry := regexp.MustCompile(`^(CO|MX|EC)$`)
	validType := regexp.MustCompile(`^(saving account|current account)$`)
	if !validType.MatchString(a.Type) {
		return errors.New("invalid account type")
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

	if !validCountry.MatchString(a.Owner.Country) {
		return errors.New("invalid owner's country")
	}

	return nil
}

func Credit(a *Account, amount money.Money) error {
	if a.Balance.Currency != amount.Currency {
		return errors.New("different currencies")
	}

	v := money.Add(a.Balance, amount)
	a.Balance = v
	return nil
}

func Debit(a *Account, amount money.Money) error {
	if a.Balance.Currency != amount.Currency {
		return errors.New("different currencies")
	}

	if !money.GtOrEq(a.Balance, amount) {
		return errors.New("insufficient balance")
	}

	v := money.Sub(a.Balance, amount)
	a.Balance = v
	return nil
}
