package subscriber

import (
	"log"
	"sync"

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
	wg := sync.WaitGroup{}
	wg.Add(1)

	// consume messages from the consumer in callback
	sub, _ := h.Conn.Subscribe(h.Cfg.Topic, func(msg *nats.Msg) {
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
