package metrics

type Config struct {
	Metric Metric `koanf:"metric"`
}

type Metric struct {
	Enabled bool   `koanf:"enabled"`
	Address string `koanf:"address"`
}
