#!/bin/bash

SRV_NAME="$1"
PROJECT_NAME=$(grep 'module' ../../go.mod | awk -F ' ' '{print $2}')
DIRECTORY="../../internal/$SRV_NAME"
if [ ! -d "$DIRECTORY" ]; then
    mkdir -p "$DIRECTORY/app"
    mkdir -p "$DIRECTORY/hello"
    mkdir -p "$DIRECTORY/config"
    mkdir -p "$DIRECTORY/pb"
fi

cat << EOF > $DIRECTORY/app/app.go
package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
    _ "net/http/pprof"

	"$PROJECT_NAME/internal/$SRV_NAME/hello"
	"$PROJECT_NAME/internal/$SRV_NAME/pb"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type (
	Deps struct {
		fx.In
		Config       *viper.Viper
		Log          *log.Logger
		HelloService *hello.Service
	}
)

func NewGRPCServer(lc fx.Lifecycle, deps Deps) *grpc.Server {
	var tracerProvider *trace.TracerProvider
	var grpcServer *grpc.Server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpointURL(deps.Config.GetString("otel.endpoint")))
			if err != nil {
				deps.Log.Fatalf("initializing otel trace grpc: %v", err)
			}

			tracerProvider = trace.NewTracerProvider(trace.WithBatcher(exp))
			otel.SetTracerProvider(tracerProvider)
			lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", deps.Config.GetInt("grpc.port")))
			if err != nil {
				return err
			}

			grpcServer = grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
			pb.RegisterHelloServiceServer(grpcServer, deps.HelloService)
			go func() {
				err := grpcServer.Serve(lis)
				if err != nil {
					deps.Log.Printf("starting grpc server: %v", err)
				}
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			if err := tracerProvider.Shutdown(ctx); err != nil {
				deps.Log.Printf("shtdown tracer provider: %v", err)
			}

			return nil
		},
	})

	return grpcServer
}

func NewHTTPServer(lc fx.Lifecycle, deps Deps) *http.Server {
	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:    deps.Config.GetString("http.addr"),
		Handler: handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			if deps.Config.GetBool("pprof") {
				go func() {
					log.Println(http.ListenAndServe(":6060", nil))
				}()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
EOF

cat << EOF > $DIRECTORY/app/module.go
package app

import (
	"$PROJECT_NAME/internal/$SRV_NAME/config"
	"$PROJECT_NAME/internal/$SRV_NAME/hello"

	"$PROJECT_NAME/pkg/kafka"
	"$PROJECT_NAME/pkg/nats"
	"$PROJECT_NAME/pkg/log"
	"$PROJECT_NAME/pkg/mongodb"

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
EOF

cat << EOF > $DIRECTORY/hello/grpc.go
package hello

import (
	"context"
	"log"
    "time"

	"$PROJECT_NAME/internal/$SRV_NAME/pb"
	"$PROJECT_NAME/pkg/kafka"
	pkglog "$PROJECT_NAME/pkg/log"
	"$PROJECT_NAME/pkg/otel"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"google.golang.org/protobuf/proto"
)

type (
	Deps struct {
		fx.In
		Log   *log.Logger
		Kafka *kafka.Conn
	}

	Service struct {
		pb.UnimplementedHelloServiceServer
		log   *log.Logger
		kafka *kafka.Conn
    }
)

const (
    v1TestedTopic = "$SRV_NAME.v1.tested"
)

func NewGRPC(deps Deps) *Service {
	s := &Service{
		log:   deps.Log,
		kafka: deps.Kafka,
	}

	deps.Kafka.Subscribe(v1TestedTopic, s.OnTested)
    return s
}

func (s *Service) World(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	evt, err := proto.Marshal(&pb.Events_V1_Tested{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
        Attributes: &pb.Events_V1_Tested_Attributes{
            Msg:        req.Msg,
        },
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
	}

	err = s.kafka.Publish(v1TestedTopic, evt)
	if err != nil {
		s.log.Printf("publishing event: %v", err)
	}

    s.log.Println(pkglog.Info("hello world!"))
    return &pb.HelloWorldResponse{Msg: req.Msg}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Start(context.Background(), otel.NoTracer, "$SRV_NAME.HelloService/OnTested")
	defer span.End()

	var evt pb.Events_V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.log.Printf("unmarshaling event: %v", err)
		return
	}

	s.log.Printf(pkglog.Info(\`msg received: {"id": %s, "occurred_on": %s, "msg": %s}\`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Attributes.Msg)
}
EOF

cat << EOF > $DIRECTORY/config/config.go
package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func New() *viper.Viper {
	v := viper.New()
    v.AddConfigPath("./configs")
    v.SetConfigName("$SRV_NAME")
	v.SetConfigType("yml")
    err := v.ReadInConfig()
    if err != nil {
        log.Fatalf("cannot read config file: %v", err)
    }

	configURL, ok := os.LookupEnv("CONFIG_URL")
	if !ok {
		return v
	}

	v.AddRemoteProvider("etcd3", configURL, "/configs/$SRV_NAME.yml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatalf("cannot read remote config: %v", err)
	}

	return v
}
EOF

exit 0
