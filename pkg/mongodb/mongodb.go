package mongodb

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type Deps struct {
	fx.In
	Config *viper.Viper
	Logger *log.Logger
}

func New(deps Deps) *mongo.Client {
	opts := options.Client().ApplyURI(deps.Config.GetString("mongodb.uri"))
	opts.Monitor = otelmongo.NewMonitor()
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		deps.Logger.Fatalf("connecting mongo: %v", err)
	}

	return client
}
