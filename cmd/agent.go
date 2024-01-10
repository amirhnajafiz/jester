package cmd

import (
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
	h := http.Handler{
		Metrics: metrics.New(a.Cfg.Metrics),
	}

	// register http handler
	if err := h.Register(a.Cfg.HTTP.Port); err != nil {
		panic(err)
	}
}
