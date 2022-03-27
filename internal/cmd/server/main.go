package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/profiling/proto"
)

type Server struct {
}

func NewServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterProfilingServer(s, &Server{})

}
