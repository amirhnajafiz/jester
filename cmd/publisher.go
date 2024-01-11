package cmd

import (
	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/publisher"
	"time"

	"github.com/spf13/cobra"
)

type Publisher struct {
	Cfg config.Config
}

func (p Publisher) Command() *cobra.Command {
	return nil
}

func (p Publisher) main() {
	h := publisher.New(publisher.Config{
		Agent:    p.Cfg.HTTP.Agent,
		Topic:    p.Cfg.NATS.Topic,
		Host:     p.Cfg.NATS.Host,
		Interval: time.Duration(p.Cfg.PublisherInterval) * time.Second,
	})

	// start handler
	err := h.Start()
	if err != nil {
		panic(err)
	}
}
