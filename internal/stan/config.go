package stan

type Config struct {
	ClusterId string `koanf:"cluster_id"`
	ClientId  string `koanf:"client_id"`
}
