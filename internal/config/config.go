package config

type Config struct {
}

type DatabaseConfig struct {
}

type ServerConfig struct{}

func Load() *Config {
	return &Config{}
}
