package config

import (
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"
)

func Default() Config {
	return Config{
		HTTPPort: 8080,
		Metrics: metrics.Config{
			Address: ":8081",
			Enabled: true,
		},
	}
}
