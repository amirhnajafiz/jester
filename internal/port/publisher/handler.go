package publisher

import (
	"log"
	"time"

	"github.com/amirhnajafiz/jester/internal/client/http"
	internalNATS "github.com/amirhnajafiz/jester/internal/client/nats"
)

type Handler struct {
	Cfg    Config
	Client http.Client
	NATS   internalNATS.Client
}

func New(cfg Config) *Handler {
	return &Handler{
		Cfg: cfg,
		Client: http.Client{
			Host: cfg.Agent,
		},
		NATS: internalNATS.Client{
			Host: cfg.Host,
		},
	}
}

func (h Handler) Start() error {
	for {
		if _, err := h.NATS.JS.Publish(h.Cfg.Topic, []byte("testing message"), nil); err != nil {
			log.Println(err)
		}

		time.Sleep(h.Cfg.Interval)
	}
}
