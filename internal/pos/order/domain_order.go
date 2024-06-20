package order

import (
	"time"

	"github.com/joseluis8906/project-layout/pkg/money"
)

type (
	Order struct {
		Number uint64    `bson:"number"`
		Date   time.Time `bson:"date"`
		Items  []Item    `bson:"items"`
	}

	Item struct {
		Name  string      `bson:"name"`
		Price money.Money `bson:"price"`
	}
)
