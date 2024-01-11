package publisher

import (
	"log"
	"time"

	"github.com/amirhnajafiz/jester/internal/client/http"
	internalNATS "github.com/amirhnajafiz/jester/internal/client/nats"
	"github.com/amirhnajafiz/jester/pkg"
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
		c.SendPost(pkg.NewRequest(pkg.FieldAddPublisher).WithLabel(cfg.Topic).ToBytes())
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
	for {
		if _, err := h.NATS.JS.Publish(h.Cfg.Topic, []byte("testing message"), nil); err != nil {
			log.Println(err)

			h.Client.SendPost(pkg.NewRequest(pkg.FieldFailures).WithLabel(h.Cfg.Topic).ToBytes())
			h.Client.SendPost(pkg.NewRequest(pkg.FieldRemovePublisher).WithLabel(h.Cfg.Topic).ToBytes())

			return err
		}

		h.Client.SendPost(pkg.NewRequest(pkg.FieldPublish).WithLabel(h.Cfg.Topic).ToBytes())

		time.Sleep(h.Cfg.Interval)
	}
}
