package publisher

import "github.com/amirhnajafiz/jester/internal/client"

type Handler struct {
	Cfg    Config
	Client client.Client
}
