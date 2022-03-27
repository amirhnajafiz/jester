package server

import (
	"log"
	"net"

	"github.com/amirhnajafiz/Stan-Gee/internal/http/handler"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"google.golang.org/grpc"
)

func NewServer(port string) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterStanGServer(s, &handler.Handler{})

	return s, lis
}
