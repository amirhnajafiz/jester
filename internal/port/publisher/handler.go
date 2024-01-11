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
	n := internalNATS.Client{
		Host: cfg.Host,
	}

	retry := 0

	for i := 0; i < cfg.MaxRetry; i++ {
		if err := n.Connect(); err == nil {
			break
		}

		retry++

		time.Sleep(5 * time.Second)
	}

	return &Handler{
		Cfg: cfg,
		Client: http.Client{
			Host: cfg.Agent,
		},
		NATS: n,
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
