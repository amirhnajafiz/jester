package publisher

import (
	"context"
	"log"
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

	for {
		if _, err := h.Conn.Publish(ctx, h.Cfg.Stream, []byte(""), nil); err != nil {
			log.Println(err)
		}

		time.Sleep(h.Cfg.Interval)
	}
}
