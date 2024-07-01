package account

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

type (
	RepoDeps struct {
		fx.In
		Mongodb *mongo.Client
	}

	Repository struct {
		db *mongo.Collection
	}
)

func NewRepository(deps RepoDeps) *Repository {
	db := deps.Mongodb.Database("mtx").Collection("accounts")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "phone_number", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &Repository{
		db: db,
	}
}

func (r *Repository) Persist(ctx context.Context, account Account) error {
	_, span := otel.Tracer("").Start(ctx, "mtx.AccountRepository/Persist")
	defer span.End()

	filter := bson.D{
		{Key: "phone_number", Value: account.PhoneNumber},
	}
	_, err := r.db.ReplaceOne(ctx, filter, account, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("replacing one account: %w", err)
	}

	return nil
}
