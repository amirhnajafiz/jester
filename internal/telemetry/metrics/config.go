package metrics

type Config struct {
	Enabled   bool   `koanf:"enabled"`
	Address   string `koanf:"address"`
	Subsystem string `koanf:"subsystem"`
	Namespace string `koanf:"namespace"`
}
