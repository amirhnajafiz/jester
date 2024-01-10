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
	// open nats connection
	conn, err := NewNATSConn(p.Cfg.NATS.Host)
	if err != nil {
		panic(err)
	}

	// register handler
	publisher.New(publisher.Config{
		Agent:    p.Cfg.HTTP.Agent,
		Topic:    p.Cfg.NATS.Topic,
		Interval: time.Duration(p.Cfg.PublisherInterval) * time.Second,
	}, conn)
}
