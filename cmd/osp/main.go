package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/joseluis8906/project-layout/internal/osp/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fx.New(
		fx.Options(app.InfraModule),
		fx.Options(app.RepoModule),
		fx.Options(app.WorkerModule),
		fx.Options(app.GRPCModule),
		fx.Provide(app.NewGRPCServer),
		fx.Provide(app.NewHTTPServer),
		fx.Invoke(func(*grpc.Server) {}),
		fx.Invoke(func(*http.Server) {}),
	)

	err := app.Start(ctx)
	if err != nil {
		log.Fatalf("stating fx app: %v", err)
	}

	<-ctx.Done()

	err = app.Stop(context.TODO())
	if err != nil {
		log.Fatalf("stoping fx app: %v", err)
	}
}
