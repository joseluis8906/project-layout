package tx

import (
	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/pkg/money"
)

type (
	Tx struct {
		ID      string      `bson:"id"`
		SrcBank Bank        `bson:"src_bank"`
		DstBank Bank        `bson:"dst_bank"`
		Amount  money.Money `bson:"amount"`
		Status  string      `bson:"status"`
	}

	Bank struct {
		Name    string          `bson:"name"`
		Account account.Account `bson:"account"`
	}
)

func (t Tx) IsZero() bool {
	return t.ID == ""
}
