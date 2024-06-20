package nats

import (
	"log"

	"github.com/nats-io/nats.go"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Deps struct {
	fx.In
	Config *viper.Viper
	Logger *log.Logger
}

func New(deps Deps) *nats.Conn {
	nc, err := nats.Connect(deps.Config.GetString("nats.url"))
	if err != nil {
		deps.Logger.Fatalf("connecting nats: %v", err)
	}

	return nc
}
