package subscriber

import (
	"time"
)

type Config struct {
	Interval time.Duration
	Agent    string
	Topic    string
	Stream   string
}
