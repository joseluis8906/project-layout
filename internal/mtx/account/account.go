package account

import (
	"errors"
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
		return errors.New("invalid phone")
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
