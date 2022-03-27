package config

import (
	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
	"github.com/amirhnajafiz/Stan-Gee/internal/stan"
	"github.com/amirhnajafiz/Stan-Gee/internal/telemetry"
)

func Default() Config {
	return Config{
		Server: server.Config{
			Type: "tcp",
			Port: ":8080",
			Stan: stan.Config{
				ClusterId: "",
				ClientId:  "",
			},
		},
		Telemetry: telemetry.Config{
			Trace: telemetry.Trace{
				Enabled: false,
				Ratio:   0.1,
				Agent: telemetry.Agent{
					Host: "127.0.0.1",
					Port: "6831",
				},
			},
			Metric: telemetry.Metric{
				Address: ":8080",
				Enabled: true,
			},
		},
	}
}
