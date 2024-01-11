package subscriber

import (
	"time"
)

type Config struct {
	Interval time.Duration
	MaxRetry int
	Host     string
	Agent    string
	Topic    string
}
