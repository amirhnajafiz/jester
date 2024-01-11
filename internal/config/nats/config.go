package nats

type Config struct {
	Host     string `koanf:"host"`
	Topic    string `koanf:"topic"`
	MaxRetry int    `koanf:"max_retry"`
}
