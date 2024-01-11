package subscriber

import (
	"time"
)

type Config struct {
	Interval time.Duration
	Host     string
	Agent    string
	Topic    string
}
