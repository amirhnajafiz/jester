package nats

import "github.com/nats-io/nats.go"

type Client struct {
	Host string
	JS   nats.JetStream
}

func (c *Client) Connect() error {
	// connect to nats server
	nc, err := nats.Connect(c.Host)
	if err != nil {
		return err
	}

	// create jetstream context from nats connection
	js, err := nc.JetStream(nil)
	if err != nil {
		return err
	}

	// keep connection
	c.JS = js

	return nil
}
