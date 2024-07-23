package app

import (
	"github.com/joseluis8906/project-layout/internal/itau/account"
	"github.com/joseluis8906/project-layout/internal/itau/config"
	"github.com/joseluis8906/project-layout/internal/itau/tx"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/mongodb"
	"github.com/joseluis8906/project-layout/pkg/nats"
	"github.com/joseluis8906/project-layout/pkg/rabbitmq"

	"go.uber.org/fx"
)

var (
	InfraModule = fx.Provide(
		config.New,
		log.New,
		mongodb.New,
		kafka.New,
		nats.New,
		rabbitmq.New,
	)

	RepoModule = fx.Provide(
		account.NewRepository,
		tx.NewRepository,
	)

	WorkerModule = fx.Provide(
		tx.NewWorker,
	)

	GRPCModule = fx.Provide(
		account.NewGRPC,
		tx.NewGRPC,
	)
)
