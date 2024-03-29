package cmd

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/publisher"

	"github.com/spf13/cobra"
)

// Publisher sends data over nats
// in an interval
type Publisher struct {
	Cfg config.Config
}

func (p Publisher) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "publisher",
		Short: "nats-publisher",
		Long:  "publisher sends data over NATS cluster",
		Run:   p.main,
	}
}

func (p Publisher) main(_ *cobra.Command, _ []string) {
	h := publisher.New(publisher.Config{
		Agent:    fmt.Sprintf("%s/cover", p.Cfg.HTTP.Agent),
		Topic:    p.Cfg.NATS.Topic,
		Host:     p.Cfg.NATS.Host,
		MaxRetry: p.Cfg.NATS.MaxRetry,
		Interval: time.Duration(p.Cfg.PublisherInterval) * time.Second,
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
