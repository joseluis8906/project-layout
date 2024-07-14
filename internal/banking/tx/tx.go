package tx

import (
	"errors"

	"github.com/joseluis8906/project-layout/pkg/money"
)

type (
	Tx struct {
		ID         string      `bson:"id"`
		SrcAccount Account     `bson:"src_account"`
		DstAccount Account     `bson:"dst_account"`
		Amount     money.Money `bson:"amount"`
		Status     string      `bson:"status"`
	}

	Account struct {
		Bank   string `bson:"bank"`
		Type   string `bson:"type"`
		Number string `bson:"number"`
	}
)

func (t Tx) IsZero() bool {
	return t.ID == ""
}

func Validate(t Tx) error {
	sameAccount := t.SrcAccount.Bank == t.DstAccount.Bank &&
		t.SrcAccount.Type == t.DstAccount.Type &&
		t.SrcAccount.Number == t.DstAccount.Number
	if sameAccount {
		return errors.New("src and dst account are the same")
	}

	return nil
}
