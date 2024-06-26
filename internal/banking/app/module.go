package app

import (
	"github.com/joseluis8906/project-layout/internal/banking/account"
	"github.com/joseluis8906/project-layout/internal/banking/config"
	"github.com/joseluis8906/project-layout/internal/banking/tx"

	"github.com/joseluis8906/project-layout/pkg/kafka"
	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/joseluis8906/project-layout/pkg/mongodb"
	"github.com/joseluis8906/project-layout/pkg/nats"
	"github.com/joseluis8906/project-layout/pkg/rabbitmq"

	"go.uber.org/fx"
)

// Module exports the module for app.
var Module = fx.Provide(
	// infra
	config.New,
	log.New,
	mongodb.New,
	kafka.New,
	nats.New,
	rabbitmq.New,

	// services
	account.New,
	tx.New,
)
