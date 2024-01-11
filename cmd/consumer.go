package cmd

import (
	"fmt"
	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/subscriber"

	"github.com/spf13/cobra"
)

type Consumer struct {
	Cfg config.Config
}

func (c Consumer) Command() *cobra.Command {
	return &cobra.Command{
		Use: "consumer",
		Run: c.main,
	}
}

func (c Consumer) main(_ *cobra.Command, _ []string) {
	h := subscriber.New(subscriber.Config{
		Agent:    fmt.Sprintf("%s/cover", c.Cfg.HTTP.Agent),
		Topic:    c.Cfg.NATS.Topic,
		Host:     c.Cfg.NATS.Host,
		MaxRetry: c.Cfg.NATS.MaxRetry,
	})

	// start handler
	err := h.Start()
	if err != nil {
		panic(err)
	}
}
