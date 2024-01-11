package nats

import "github.com/nats-io/nats.go"

type Client struct {
	Host string
}

func (c Client) Connect() (nats.JetStream, error) {
	// connect to nats server
	nc, err := nats.Connect(c.Host)
	if err != nil {
		return nil, err
	}

	// create jetstream context from nats connection
	return nc.JetStream(nil)
}
