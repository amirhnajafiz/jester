package handler

import (
	"context"
	"log"

	"github.com/amirhnajafiz/Stan-Gee/proto"
	"github.com/nats-io/stan.go"
)

type Handler struct {
	Stan stan.Conn
}

func (h *Handler) Sub(in *proto.Send, stream proto.StanG_SubServer) error {
	// send messages via nats
	sub, _ := h.Stan.Subscribe(in.Topic, func(msg *stan.Msg) {
		err := stream.Send(&proto.Catch{
			Content: string(msg.Data),
		})
		if err != nil {
			log.Fatalf("faield to subscribe: %v\n", err)
		}
	})

	defer func(sub stan.Subscription) {
		_ = sub.Unsubscribe()
	}(sub)

	return nil
}

func (h *Handler) Put(_ context.Context, in *proto.Data) (*proto.Response, error) {
	// get message from nats
	err := h.Stan.Publish(in.Topic, []byte(in.Content))
	if err != nil {
		return &proto.Response{
			Status:  -1,
			Message: err.Error(),
		}, err
	}

	return &proto.Response{
		Status:  1,
		Message: "Successfully sent",
	}, nil
}
