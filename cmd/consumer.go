package cmd

import (
	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/subscriber"

	"github.com/spf13/cobra"
)

type Consumer struct {
	Cfg config.Config
}

func (c Consumer) Command() *cobra.Command {
	return nil
}

func (c Consumer) main() {
	// register handler
	err := subscriber.New(subscriber.Config{
		Agent: c.Cfg.HTTP.Agent,
		Topic: c.Cfg.NATS.Topic,
		Host:  c.Cfg.NATS.Host,
	}).Start()
	if err != nil {
		panic(err)
	}
}
