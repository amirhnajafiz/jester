package subscriber

import (
	"log"
	"sync"

	"github.com/amirhnajafiz/jester/internal/client/http"
	internalNATS "github.com/amirhnajafiz/jester/internal/client/nats"

	"github.com/nats-io/nats.go"
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

func (h Handler) Register() error {
	return h.NATS.Connect()
}

func (h Handler) Start() error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// consume messages from the consumer in callback
	sub, _ := h.NATS.JS.Subscribe(h.Cfg.Topic, func(msg *nats.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)
		}

		log.Println("received jetstream message: ", string(msg.Data))
	})
	defer func(sub *nats.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			log.Println(err)
		}
	}(sub)

	wg.Wait()

	return nil
}
