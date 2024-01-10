package config

import (
	"log"

	"github.com/amirhnajafiz/jester/internal/config/http"
	"github.com/amirhnajafiz/jester/internal/config/nats"
	"github.com/amirhnajafiz/jester/internal/telemetry/metrics"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	HTTP              http.Config    `koanf:"http"`
	NATS              nats.Config    `koanf:"nats"`
	Metrics           metrics.Config `koanf:"metrics"`
	PublisherInterval int            `koanf:"publisher_interval"`
}

func Load(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Fatalf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
