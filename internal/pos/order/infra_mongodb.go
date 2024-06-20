package order

import "go.mongodb.org/mongo-driver/mongo"

type (
	Repository struct {
		db *mongo.Database
	}
)
