package stan

import (
	"log"

	"github.com/nats-io/stan.go"
)

func Connect(cfg Config) stan.Conn {
	sc, err := stan.Connect(cfg.ClusterId, cfg.ClientId)
	if err != nil {
		log.Fatalf("failed to connect to nats-stream server: %v\n", err)
	}

	return sc
}
