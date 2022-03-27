package server

import "github.com/amirhnajafiz/Stan-Gee/internal/stan"

type Config struct {
	Type string      `koanf:"type"`
	Port string      `koanf:"port"`
	Stan stan.Config `koanf:"stan"`
}
