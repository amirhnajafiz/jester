package cmd

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func NewNATSConn(host string) (jetstream.JetStream, error) {
	// connect to nats server
	nc, err := nats.Connect(host)
	if err != nil {
		return nil, err
	}

	// create jetstream context from nats connection
	return jetstream.New(nc)
}
