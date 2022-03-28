package server

import (
	"log"
	"net"

	"github.com/amirhnajafiz/Stan-Gee/internal/http/handler"
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func NewServer(cfg Config, trace trace.Tracer, metric stan.Metrics) (*grpc.Server, net.Listener) {
	lis, err := net.Listen(cfg.Type, cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	c := stan.Connect(cfg.Stan)
	if c == nil {
		metric.ConnectionErrors.Add(1)
	}

	s := grpc.NewServer()
	proto.RegisterStanGServer(s, &handler.Handler{
		Tracer:  trace,
		Stan:    c,
		Metrics: metric,
	})

	return s, lis
}
