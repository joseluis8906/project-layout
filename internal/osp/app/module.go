package app

import (
	"github.com/joseluis8906/project-layout/internal/osp/config"
	"github.com/joseluis8906/project-layout/internal/osp/hello"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/mongodb"
	"github.com/joseluis8906/project-layout/pkg/nats"

	"go.uber.org/fx"
)

var (
	InfraModule = fx.Provide(
		config.New,
		log.New,
		mongodb.New,
		kafka.New,
		nats.New,
	)

	RepoModule = fx.Provide()

	WorkerModule = fx.Provide()

	GRPCModule = fx.Provide(
		hello.NewGRPC,
	)
)
