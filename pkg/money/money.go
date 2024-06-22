package money

import "github.com/Rhymond/go-money"

const (
	USD = money.USD
	COP = money.COP
	MXN = money.MXN
)

type Money struct {
	money    *money.Money `bson:"-"`
	Amount   int64        `bson:"amount"`
	Currency string       `bson:"currency"`
}

func New(amount int64, currency string) Money {
	m := money.New(amount, currency)
	return Money{money: m, Amount: m.Amount(), Currency: m.Currency().Code}
}

func (m Money) IsZero() bool {
	return m.Currency == ""
}

func (m Money) Add(a Money) Money {
	if m.money == nil {
		m.money = money.New(m.Amount, m.Currency)
	}

	n, err := m.money.Add(money.New(a.Amount, a.Currency))
	if err != nil {
		return Money{}
	}

	return Money{money: n, Amount: n.Amount(), Currency: n.Currency().Code}
}

func (m Money) String() string {
	if m.money == nil {
		return "$0"
	}

	return m.money.Display()
}
