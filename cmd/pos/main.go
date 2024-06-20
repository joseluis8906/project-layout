package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/joseluis8906/project-layout/internal/pos/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fx.New(
		fx.Options(app.Module),
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
