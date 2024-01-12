package cmd

import (
	"github.com/amirhnajafiz/jester/internal/cache"
	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/http"
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"

	"github.com/spf13/cobra"
)

// Agent is responsible for getting publisher and consumer
// metrics and expose them over an endpoint
type Agent struct {
	Cfg config.Config
}

func (a Agent) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "agent",
		Short: "metrics-server",
		Long:  "agent is the Jester metrics server",
		Run:   a.main,
	}
}

func (a Agent) main(_ *cobra.Command, _ []string) {
	// register metrics
	m, err := metrics.New(a.Cfg.Metrics)
	if err != nil {
		panic(err)
	}

	// open etcd connection
	e, err := cache.NewCache(a.Cfg.ETCD)
	if err != nil {
		panic(err)
	}

	// register handler
	h := http.Handler{
		ETCD:    e,
		Metrics: m,
	}

	// start handler
	if err := h.Register(a.Cfg.HTTP.Port, a.Cfg.Metrics.Enabled); err != nil {
		panic(err)
	}
}
