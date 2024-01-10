package publisher

import (
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
