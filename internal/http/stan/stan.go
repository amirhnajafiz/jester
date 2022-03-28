package stan

import (
	"github.com/nats-io/stan.go"
)

func Connect(cfg Config) (stan.Conn, error) {
	return stan.Connect(cfg.ClusterId, cfg.ClientId)
}
