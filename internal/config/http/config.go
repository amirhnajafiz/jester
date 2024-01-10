package http

type Config struct {
	Agent string `koanf:"agent"`
	Port  int    `koanf:"port"`
}
