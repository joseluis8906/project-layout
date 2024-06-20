package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In
	Config *viper.Viper
	Logger *log.Logger
}

func New(deps Deps) *amqp.Connection {
	conn, err := amqp.Dial(deps.Config.GetString("rabbitmq.url"))
	if err != nil {
		deps.Logger.Fatalf("connecting rabbitmq: %v", err)
	}

	return conn
}
