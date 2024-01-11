package cmd

import (
	"github.com/amirhnajafiz/jester/internal/cache"
	"github.com/amirhnajafiz/jester/internal/config"
	"github.com/amirhnajafiz/jester/internal/port/http"
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"

	"github.com/spf13/cobra"
)

type Agent struct {
	Cfg config.Config
}

func (a Agent) Command() *cobra.Command {
	return &cobra.Command{
		Use: "agent",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a Agent) main() {
	// start metrics server
	metrics.NewServer(a.Cfg.Metrics).Start()

	// register metrics
	m, err := metrics.New(a.Cfg.Metrics)
	if err != nil {
		panic(err)
	}

	// etcd
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
	if err := h.Register(a.Cfg.HTTP.Port); err != nil {
		panic(err)
	}
}
