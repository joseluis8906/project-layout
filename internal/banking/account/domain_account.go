package account

import (
	"errors"
	"regexp"
)

type (
	Account struct {
		Type   string `bson:"type"`
		Number string `bson:"number"`
		Owner  Owner  `bson:"owner"`
	}

	Owner struct {
		ID       string `bson:"id"`
		Country  string `bson:"country"`
		Email    string `bson:"email"`
		FullName string `bson:"full_name"`
	}
)

func (a Account) Validate() error {
	validID := regexp.MustCompile(`^[0-9]+$`)
	validEmail := regexp.MustCompile(`^[\w]+.*@[\w]+.(com|net|org)`)
	validName := regexp.MustCompile(`^[\w]{2,} [\w]{2,}$`)
	validCountry := regexp.MustCompile(`^(CO|MX|EC)$`)
	if !validID.MatchString(a.Owner.ID) {
		return errors.New("invalid id")
	}

	if !validEmail.MatchString(a.Owner.Email) {
		return errors.New("invalid email")
	}

	if !validName.MatchString(a.Owner.FullName) {
		return errors.New("invalid name")
	}

	if !validCountry.MatchString(a.Owner.Country) {
		return errors.New("invalid country")
	}

	return nil
}
