package account

import (
	"context"
	"fmt"

	"github.com/joseluis8906/project-layout/pkg/otel"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type (
	RepoDeps struct {
		fx.In
		Conf    *viper.Viper
		Mongodb *mongo.Client
	}

	Repository struct {
		db *mongo.Collection
	}
)

func NewRepository(deps RepoDeps) *Repository {
	db := deps.Mongodb.Database(deps.Conf.GetString("app")).Collection("accounts")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "type", Value: -1}, {Key: "number", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &Repository{
		db: db,
	}
}

func (r *Repository) Persist(ctx context.Context, account Account) error {
	_, span := otel.Start(ctx, otel.NoTracer, "bank.AccountRepository/Persist")
	defer span.End()

	filter := bson.D{
		{Key: "type", Value: account.Type},
		{Key: "number", Value: account.Number},
	}
	_, err := r.db.ReplaceOne(ctx, filter, account, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("replacing one document: %w", err)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, aType, number string) (Account, error) {
	_, span := otel.Start(ctx, otel.NoTracer, "bank.AccountRepository/Get")
	defer span.End()

	filter := bson.D{
		{Key: "type", Value: aType},
		{Key: "number", Value: number},
	}
	cur := r.db.FindOne(ctx, filter)
	if err := cur.Err(); err != nil {
		return Account{}, fmt.Errorf("finding one document: %w", err)
	}

	var account Account
	if err := cur.Decode(&account); err != nil {
		return Account{}, fmt.Errorf("decoding account: %w", err)
	}

	return account, nil
}
