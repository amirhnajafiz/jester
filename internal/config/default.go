package config

import (
	"github.com/amirhnajafiz/Stan-Gee/internal/cmd/server"
	"github.com/amirhnajafiz/Stan-Gee/internal/stan"
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
	}
}
