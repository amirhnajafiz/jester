package cmd

import "github.com/nats-io/nats.go"

func NewNATSConn(host string) (nats.JetStream, error) {
	// connect to nats server
	nc, err := nats.Connect(host)
	if err != nil {
		return nil, err
	}

	// create jetstream context from nats connection
	return nc.JetStream(nil)
}
