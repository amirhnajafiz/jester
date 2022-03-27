package server

import (
	"log"
	"net"

	"github.com/amirhnajafiz/Stan-Gee/internal/http/handler"
	"github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"google.golang.org/grpc"
)

func NewServer(cfg Config) (*grpc.Server, net.Listener) {
	lis, err := net.Listen(cfg.Type, cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterStanGServer(s, &handler.Handler{Stan: stan.Connect(cfg.Stan)})

	return s, lis
}
