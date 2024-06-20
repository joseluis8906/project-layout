package money

type Money struct {
	Amount   uint64 `bson:"amount"`
	Currency string `bson:"currency"`
	Decimals uint8  `bson:"decimals"`
}
