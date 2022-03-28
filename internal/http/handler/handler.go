package handler

import (
	"context"
	"log"

	stanMetric "github.com/amirhnajafiz/Stan-Gee/internal/http/stan"
	"github.com/amirhnajafiz/Stan-Gee/proto"
	"github.com/nats-io/stan.go"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	Metrics stanMetric.Metrics
	Tracer  trace.Tracer
	Stan    stan.Conn
}

func (h *Handler) Sub(in *proto.Send, stream proto.StanG_SubServer) error {
	_, span := h.Tracer.Start(context.Background(), "http.handler.subscribe")
	defer span.End()

	// send messages via nats
	sub, _ := h.Stan.Subscribe(in.Topic, func(msg *stan.Msg) {
		err := stream.Send(&proto.Catch{
			Content: string(msg.Data),
		})
		if err != nil {
			h.Metrics.FailedSub.Add(1)

			log.Fatalf("faield to subscribe: %v\n", err)
		} else {
			h.Metrics.SuccessfulSub.Add(1)
		}
	})

	defer func(sub stan.Subscription) {
		_ = sub.Unsubscribe()
	}(sub)

	return nil
}

func (h *Handler) Put(_ context.Context, in *proto.Data) (*proto.Response, error) {
	_, span := h.Tracer.Start(context.Background(), "http.handler.publish")
	defer span.End()

	// get message from nats
	err := h.Stan.Publish(in.Topic, []byte(in.Content))
	if err != nil {
		h.Metrics.FailedPut.Add(1)

		return &proto.Response{
			Status:  -1,
			Message: err.Error(),
		}, err
	}

	h.Metrics.SuccessfulPut.Add(1)

	return &proto.Response{
		Status:  1,
		Message: "Successfully sent",
	}, nil
}
