package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/subscriber"

	"github.com/spf13/cobra"
)

// Consumer is the NATS subscriber
type Consumer struct {
	Cfg config.Config
}

func (c Consumer) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "nats-subscriber",
		Long:  "consumer subscribes on a topic over NATS",
		Run:   c.main,
	}
}

func (c Consumer) main(_ *cobra.Command, _ []string) {
	h := subscriber.New(subscriber.Config{
		Agent:    fmt.Sprintf("%s/cover", c.Cfg.HTTP.Agent),
		Topic:    c.Cfg.NATS.Topic,
		Host:     c.Cfg.NATS.Host,
		MaxRetry: c.Cfg.NATS.MaxRetry,
	})
	if h == nil {
		panic("cannot connect to NATS cluster")
	}

	// start handler
	err := h.Start()
	if err != nil {
		panic(err)
	}
}
