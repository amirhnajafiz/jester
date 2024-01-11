package config

import (
	"github.com/amirhnajafiz/jester/internal/cache"
	"github.com/amirhnajafiz/jester/internal/config/http"
	"github.com/amirhnajafiz/jester/internal/config/nats"
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"
)

func Default() Config {
	return Config{
		ETCD: cache.Config{},
		HTTP: http.Config{
			Port:  8080,
			Agent: "",
		},
		NATS: nats.Config{},
		Metrics: metrics.Config{
			Address:   ":8081",
			Enabled:   true,
			Subsystem: "jester",
			Namespace: "snappcloud.io",
		},
		PublisherInterval: 10, // in seconds
	}
}
