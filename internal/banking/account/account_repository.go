package account

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
)

type (
	AccountRepo struct {
		db *mongo.Collection
	}
)

func NewAccountRepo(conn *mongo.Client) *AccountRepo {
	db := conn.Database("banking").Collection("accounts")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "type", Value: -1}, {Key: "number", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &AccountRepo{
		db: db,
	}
}

func (r *AccountRepo) Persist(ctx context.Context, account Account) error {
	_, span := otel.Tracer("").Start(ctx, "banking.AccountRepository/Add")
	defer span.End()

	filter := bson.D{{Key: "type", Value: account.Type}, {Key: "number", Value: account.Number}}
	_, err := r.db.ReplaceOne(ctx, filter, account, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("upserting account: %w", err)
	}

	return nil
}

func (r *AccountRepo) Get(ctx context.Context, aType string, number string) (Account, error) {
	_, span := otel.Tracer("").Start(ctx, "banking.AccountRepository/Get")
	defer span.End()

	filter := bson.D{{Key: "type", Value: aType}, {Key: "number", Value: number}}
	cur := r.db.FindOne(ctx, filter)
	if err := cur.Err(); err != nil {
		return Account{}, fmt.Errorf("upserting account: %w", err)
	}

	var account Account
	if err := cur.Decode(&account); err != nil {
		return Account{}, fmt.Errorf("decoding account: %w", err)
	}

	return account, nil
}
