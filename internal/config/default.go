package config

import (
	"github.com/amirhnajafiz/jester/internal/config/http"
	"github.com/amirhnajafiz/jester/internal/config/nats"
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"
)

func Default() Config {
	return Config{
		HTTP: http.Config{
			Port: 8080,
		},
		NATS: nats.Config{},
		Metrics: metrics.Config{
			Address:   ":8081",
			Enabled:   true,
			Subsystem: "jester",
			Namespace: "snappcloud.io",
		},
	}
}
