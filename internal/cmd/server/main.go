package server

import (
	"context"
	"log"
	"net"

	"github.com/amirhnajafiz/Stan-Gee/internal/http/handler"
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func NewServer(cfg Config, trace trace.Tracer, metric stan.Metrics) (*grpc.Server, net.Listener) {
	_, span := trace.Start(context.Background(), "server.new.server")
	defer span.End()

	lis, err := net.Listen(cfg.Type, cfg.Port)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	c, err := stan.Connect(cfg.Stan)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		metric.ConnectionErrors.Add(1)

		log.Fatalf("failed to connect to nats-stream server: %v\n", err)
	}

	s := grpc.NewServer()
	proto.RegisterStanGServer(s, &handler.Handler{
		Tracer:  trace,
		Stan:    c,
		Metrics: metric,
	})

	return s, lis
}
