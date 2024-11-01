package app

import (
	"github.com/joseluis8906/project-layout/internal/bank/account"
	"github.com/joseluis8906/project-layout/internal/bank/config"
	"github.com/joseluis8906/project-layout/internal/bank/tx"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/metric"
	"github.com/joseluis8906/project-layout/pkg/mongodb"
	"github.com/joseluis8906/project-layout/pkg/nats"

	"go.uber.org/fx"
)

var (
	InfraModule = fx.Provide(
		config.New,
		log.New,
		metric.New,
		mongodb.New,
		kafka.New,
		nats.New,
	)

	RepoModule = fx.Provide(
		account.NewRepository,
		tx.NewRepository,
	)

	WorkerModule = fx.Provide()

	GRPCModule = fx.Provide(
		account.NewGRPC,
		tx.NewGRPC,
	)
)
