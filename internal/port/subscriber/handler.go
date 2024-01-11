package subscriber

import (
	"log"
	"sync"
	"time"

	"github.com/amirhnajafiz/jester/internal/client/http"
	internalNATS "github.com/amirhnajafiz/jester/internal/client/nats"
	"github.com/amirhnajafiz/jester/pkg"

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

	c := http.Client{
		Host: cfg.Agent,
	}

	retry := 0

	for i := 0; i < cfg.MaxRetry; i++ {
		if err := n.Connect(); err == nil {
			break
		}

		retry++

		time.Sleep(5 * time.Second)
	}

	if retry == cfg.MaxRetry {
		c.SendPost(pkg.NewRequest(pkg.FieldFailedConnections).ToBytes())

		return nil
	} else {
		c.SendPost(pkg.NewRequest(pkg.FieldAddSubscriber).WithLabel(cfg.Topic).ToBytes())
		if retry > 0 {
			c.SendPost(pkg.NewRequest(pkg.FieldRetryPerConnection).WithValue(float64(retry)).ToBytes())
		}
	}

	return &Handler{
		Cfg:    cfg,
		Client: c,
		NATS:   n,
	}
}

func (h Handler) Start() error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// consume messages from the consumer in callback
	sub, _ := h.NATS.JS.Subscribe(h.Cfg.Topic, func(msg *nats.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)

			h.Client.SendPost(pkg.NewRequest(pkg.FieldRemoveSubscriber).WithLabel(h.Cfg.Topic).ToBytes())

			wg.Done()
		}

		h.Client.SendPost(pkg.NewRequest(pkg.FieldConsume).WithLabel(h.Cfg.Topic).WithParam(string(msg.Data)).ToBytes())

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
