package cache

type Config struct {
	Endpoints []string `koanf:"endpoints"`
	Timeout   int      `koan:"timeout"`
}
