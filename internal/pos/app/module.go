package app

import (
	"github.com/joseluis8906/project-layout/internal/pos/config"
	"github.com/joseluis8906/project-layout/internal/pos/order"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/mongodb"

	"go.uber.org/fx"
)

// Module exports the module for app.
var Module = fx.Provide(
	//infra
	config.New,
	log.New,
	mongodb.New,
	kafka.New,
	//services
	order.New,
)
