package publisher

import "time"

type Config struct {
	Host     string
	Agent    string
	Interval time.Duration
}
