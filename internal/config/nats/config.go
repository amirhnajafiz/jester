package nats

type Config struct {
	Host      string `koanf:"host"`
	Stream    string `koanf:"stream"`
	Topic     string `koanf:"topic"`
	TopicName string `koanf:"topic_name"`
}
