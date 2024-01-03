package config

import (
	"github.com/amirhnajafiz/jester/internal/telemetry"
)

func Default() Config {
	return Config{
		Telemetry: telemetry.Config{
			Metric: telemetry.Metric{
				Address: ":8080",
				Enabled: true,
			},
		},
	}
}
