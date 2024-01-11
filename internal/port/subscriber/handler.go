package subscriber

import (
	"log"
	"sync"
	"time"

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
