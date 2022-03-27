package server

import (
	"go.opentelemetry.io/otel/trace"
	"log"
	"net"

	"github.com/amirhnajafiz/Stan-Gee/internal/http/handler"
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"google.golang.org/grpc"
)

func NewServer(cfg Config, trace trace.Tracer, metric stan.Metrics) (*grpc.Server, net.Listener) {
	lis, err := net.Listen(cfg.Type, cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterStanGServer(s, &handler.Handler{
		Tracer:  trace,
		Stan:    stan.Connect(cfg.Stan),
		Metrics: metric,
	})

	return s, lis
}
