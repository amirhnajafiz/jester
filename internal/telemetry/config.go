package telemetry

type Config struct {
	Trace  Trace  `koanf:"trace"`
	Metric Metric `koanf:"metric"`
}

type Trace struct {
	Enabled bool `koanf:"enabled"`
	Agent   `koanf:"agent"`
	Ratio   float64 `koanf:"ratio"`
}

type Agent struct {
	Host string `koanf:"host"`
	Port string `koanf:"port"`
}

type Metric struct {
	Enabled bool   `koanf:"enabled"`
	Address string `koanf:"address"`
}
