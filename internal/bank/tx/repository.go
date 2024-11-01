package tx

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
		MongoDB *mongo.Client
	}

	Repository struct {
		db *mongo.Collection
	}
)

func NewRepository(deps RepoDeps) *Repository {
	db := deps.MongoDB.Database(deps.Conf.GetString("app.name")).Collection("txs")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &Repository{
		db: db,
	}
}

func (r *Repository) Persist(ctx context.Context, tx Tx) error {
	_, span := otel.Start(ctx, otel.NoTracer, "bank.TxRepository/Persist")
	defer span.End()

	filter := bson.D{{Key: "id", Value: tx.ID}}
	_, err := r.db.ReplaceOne(ctx, filter, tx, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("upserting tx: %w", err)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, id string) (Tx, error) {
	_, span := otel.Start(ctx, otel.NoTracer, "bank.TxRepository/Get")
	defer span.End()

	filter := bson.D{{Key: "id", Value: id}}
	cur := r.db.FindOne(ctx, filter)
	if err := cur.Err(); err != nil {
		return Tx{}, fmt.Errorf("finding tx: %w", err)
	}

	var tx Tx
	if err := cur.Decode(&tx); err != nil {
		return Tx{}, fmt.Errorf("decoding tx: %w", err)
	}

	return tx, nil
}
