package handler

import (
	"context"

	"github.com/amirhnajafiz/Stan-Gee/proto"
)

type Handler struct{}

func (h *Handler) Sub(in *proto.Send, stream proto.StanG_SubServer) error {
	// send messages via nats

	return nil
}

func (h *Handler) Put(ctx context.Context, in *proto.Data) (*proto.Response, error) {
	// get message from nats

	return &proto.Response{
		Status:  1,
		Message: "test",
	}, nil
}
