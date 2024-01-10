package subscriber

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/amirhnajafiz/jester/internal/client"

	"github.com/nats-io/nats.go/jetstream"
)

type Handler struct {
	Cfg    Config
	Client client.Client
	Conn   jetstream.JetStream
}

func New(cfg Config, conn jetstream.JetStream) *Handler {
	return &Handler{
		Conn: conn,
		Cfg:  cfg,
		Client: client.Client{
			Host: cfg.Agent,
		},
	}
}

func (h Handler) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// get existing stream handle
	stream, err := h.Conn.Stream(ctx, h.Cfg.Stream)
	if err != nil {
		return err
	}

	// retrieve consumer handle from a stream
	cons, err := stream.Consumer(ctx, h.Cfg.Topic)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// consume messages from the consumer in callback
	cc, _ := cons.Consume(func(msg jetstream.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)
		}

		log.Println("received jetstream message: ", string(msg.Data()))
	})
	defer cc.Stop()

	wg.Wait()

	return nil
}
