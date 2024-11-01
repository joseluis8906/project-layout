package app

import (
	"github.com/joseluis8906/project-layout/internal/mtx/account"
	"github.com/joseluis8906/project-layout/internal/mtx/config"

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
	)

	GRPCModule = fx.Provide(
		account.NewGRPC,
	)

	WorkerModule = fx.Provide()
)
