package app

import (
	"github.com/joseluis8906/project-layout/internal/itau/account"
	"github.com/joseluis8906/project-layout/internal/itau/config"
	"github.com/joseluis8906/project-layout/internal/itau/tx"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/mongodb"
	"github.com/joseluis8906/project-layout/pkg/nats"

	"go.uber.org/fx"
)

// Module exports the module for app.
var Module = fx.Provide(
	//infra
	config.New,
	log.New,
	mongodb.New,
	kafka.New,
	nats.New,

	//repositories
	account.NewRepository,
	tx.NewRepository,

	//services
	account.NewGRPC,
	tx.GRPC,
)
