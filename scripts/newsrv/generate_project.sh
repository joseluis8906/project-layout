#!/bin/bash
#

SRV_NAME="$1"
PROJECT_NAME=$(grep 'module' ../go.mod | awk -F ' ' '{print $2}')
DIRECTORY="../internal/$SRV_NAME"
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

// Module exports the module for app.
var Module = fx.Provide(
	//infra
	config.New,
	log.New,
	mongodb.New,
	kafka.New,
	nats.New,

	//services
	hello.New,
)
EOF

cat << EOF > $DIRECTORY/hello/api_grpc.go
package hello

import (
	"context"
	"log"
    "time"

	"$PROJECT_NAME/internal/$SRV_NAME/pb"
	"$PROJECT_NAME/pkg/kafka"
	loglevel "$PROJECT_NAME/pkg/log"
	evtpb "$PROJECT_NAME/pkg/pb"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
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

func New(deps Deps) *Service {
	s := &Service{
		log:   deps.Log,
		kafka: deps.Kafka,
	}

	deps.Kafka.Subscribe("v1.tested", s.OnTested)
    return s
}

func (s *Service) World(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	_, span := otel.Tracer("").Start(context.Background(), "$SRV_NAME.HelloService/World")
	defer span.End()

	evt, err := proto.Marshal(&evtpb.V1_Tested{
		Id:         uuid.New().String(),
		OccurredOn: time.Now().UnixMilli(),
        Msg:        req.Msg,
	})
	if err != nil {
		s.log.Printf("marshaling event: %v", err)
	}

	err = s.kafka.Publish("v1.tested", evt)
	if err != nil {
		s.log.Printf("publishing event: %v", err)
	}

    s.log.Println(loglevel.Info("hello world!"))
    return &pb.HelloWorldResponse{Msg: req.Msg}, nil
}

func (s *Service) OnTested(msg *kafka.Message) {
	_, span := otel.Tracer("").Start(context.Background(), "$SRV_NAME.HelloService/OnTested")
	defer span.End()

	var evt evtpb.V1_Tested
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		s.log.Printf("unmarshaling event: %v", err)
		return
	}

	s.log.Printf(loglevel.Info(\`msg received: {"id": %s, "occurred_on": %s, "msg": %s}\`), evt.Id, time.UnixMilli(evt.OccurredOn), evt.Msg)
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
	v.AddRemoteProvider("etcd3", os.Getenv("CONFIG_URL"), "/configs/$SRV_NAME.yml")
	v.SetConfigType("yml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatalf("cannot read remote config: %v", err)
	}

	return v
}
EOF

exit 0
