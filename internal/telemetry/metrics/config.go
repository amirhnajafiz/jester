package metrics

type Config struct {
	Enabled   bool   `koanf:"enabled"`
	Subsystem string `koanf:"subsystem"`
	Namespace string `koanf:"namespace"`
}
