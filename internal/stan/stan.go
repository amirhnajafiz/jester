package stan

import (
	"log"

	"github.com/nats-io/stan.go"
)

func Connect(ClusterID string, ClientID string) stan.Conn {
	sc, err := stan.Connect(ClusterID, ClientID)
	if err != nil {
		log.Fatalf("failed to connect to nats-stream server: %v\n", err)
	}

	return sc
}
