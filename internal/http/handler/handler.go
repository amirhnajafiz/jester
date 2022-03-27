package handler

import (
	"context"

	"github.com/amirhnajafiz/Stan-Gee/proto"
)

type Handler struct{}

func (h *Handler) Sub(ctx context.Context, in *proto.Send) (*proto.Catch, error) {
	// send messages via nats

	return &proto.Catch{
		Content: in.Topic,
	}, nil
}

func (h *Handler) Put(ctx context.Context, in *proto.Data) (*proto.Response, error) {
	// get message from nats

	return &proto.Response{
		Status:  1,
		Message: "test",
	}, nil
}
