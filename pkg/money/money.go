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

func (m Money) String() string {
	if m.money == nil {
		return "$0"
	}

	return m.money.Display()
}

func Add(addendr, addendl Money) Money {
	if addendr.money == nil {
		addendr.money = money.New(addendr.Amount, addendr.Currency)
	}

	n, err := addendr.money.Add(money.New(addendl.Amount, addendl.Currency))
	if err != nil {
		return Money{}
	}

	return Money{money: n, Amount: n.Amount(), Currency: n.Currency().Code}
}

func Sub(minued, subtrahend Money) Money {
	if minued.money == nil {
		minued.money = money.New(minued.Amount, minued.Currency)
	}

	res, err := minued.money.Subtract(money.New(subtrahend.Amount, subtrahend.Currency))
	if err != nil {
		return Money{}
	}

	return Money{money: res, Amount: res.Amount(), Currency: res.Currency().Code}
}

func GtOrEq(left, right Money) bool {
	if left.money == nil {
		left.money = money.New(left.Amount, left.Currency)
	}

	ok, err := left.money.GreaterThanOrEqual(money.New(right.Amount, right.Currency))
	if err != nil {
		return false
	}

	return ok
}
