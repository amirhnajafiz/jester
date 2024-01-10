package publisher

import (
	"github.com/amirhnajafiz/jester/internal/client"

	"github.com/nats-io/nats.go"
)

type Handler struct {
	Cfg    Config
	Client client.Client
	Conn   *nats.Conn
}
