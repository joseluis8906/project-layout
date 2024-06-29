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
	db := deps.Mongodb.Database("banking").Collection("accounts")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "bank", Value: -1}, {Key: "type", Value: -1}, {Key: "number", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &Repository{
		db: db,
	}
}

func (r *Repository) Persist(ctx context.Context, account Account) error {
	_, span := otel.Tracer("").Start(ctx, "banking.AccountRepository/Add")
	defer span.End()

	filter := bson.D{
		{Key: "bank", Value: account.Bank},
		{Key: "type", Value: account.Type},
		{Key: "number", Value: account.Number},
	}
	_, err := r.db.ReplaceOne(ctx, filter, account, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("upserting account: %w", err)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, bank, aType, number string) (Account, error) {
	_, span := otel.Tracer("").Start(ctx, "banking.AccountRepository/Get")
	defer span.End()

	filter := bson.D{
		{Key: "bank", Value: bank},
		{Key: "type", Value: aType},
		{Key: "number", Value: number},
	}
	cur := r.db.FindOne(ctx, filter)
	if err := cur.Err(); err != nil {
		return Account{}, fmt.Errorf("finding account: %w", err)
	}

	var account Account
	if err := cur.Decode(&account); err != nil {
		return Account{}, fmt.Errorf("decoding account: %w", err)
	}

	return account, nil
}
