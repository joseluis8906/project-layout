package tx

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
)

type (
	TxRepo struct {
		db *mongo.Collection
	}
)

func NewTxRepo(conn *mongo.Client) *TxRepo {
	db := conn.Database("banking").Collection("txs")
	db.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	return &TxRepo{
		db: db,
	}
}

func (r *TxRepo) Persist(ctx context.Context, tx Tx) error {
	_, span := otel.Tracer("").Start(ctx, "banking.TxRepository/Persist")
	defer span.End()

	filter := bson.D{{Key: "type", Value: tx.ID}}
	_, err := r.db.ReplaceOne(ctx, filter, tx, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("upserting tx: %w", err)
	}

	return nil
}

func (r *TxRepo) Get(ctx context.Context, id string) (Tx, error) {
	_, span := otel.Tracer("").Start(ctx, "banking.TxRepository/Get")
	defer span.End()

	filter := bson.D{{Key: "id", Value: id}}
	cur := r.db.FindOne(ctx, filter)
	if err := cur.Err(); err != nil {
		return Tx{}, fmt.Errorf("upserting tx: %w", err)
	}

	var tx Tx
	if err := cur.Decode(&tx); err != nil {
		return Tx{}, fmt.Errorf("decoding tx: %w", err)
	}

	return tx, nil
}
