package publisher

import (
	"log"
	"time"

	"github.com/amirhnajafiz/jester/internal/client"

	"github.com/nats-io/nats.go"
)

type Handler struct {
	Cfg    Config
	Client client.Client
	Conn   nats.JetStream
}

func New(cfg Config, conn nats.JetStream) *Handler {
	return &Handler{
		Conn: conn,
		Cfg:  cfg,
		Client: client.Client{
			Host: cfg.Agent,
		},
	}
}

func (h Handler) Start() error {
	for {
		if _, err := h.Conn.Publish(h.Cfg.Topic, []byte("testing message"), nil); err != nil {
			log.Println(err)
		}

		time.Sleep(h.Cfg.Interval)
	}
}
