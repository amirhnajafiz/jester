package handler

import (
	"context"

	"github.com/amirhnajafiz/Stan-Gee/proto"
)

func (s *server.Server) Sub(ctx context.Context, in *proto.Send) (*proto.Catch, error) {
	// send messages via nats

	return &proto.Catch{
		Content: in.Topic,
	}, nil
}

func (s *server.Server) Put(ctx context.Context, in *proto.Data) (*proto.Response, error) {
	// get message from nats

	return &proto.Response{
		Status:  1,
		Message: "test",
	}, nil
}
