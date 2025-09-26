package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	NATS   NATSConfig
	Logger Logger
}

type Logger struct {
	Level string `env:"LOG_LEVEL" env-default:"info"`
}

type NATSConfig struct {
	URL        string `env:"NATS_URL" env-default:"nats://localhost:4222"`
	ClientName string `env:"NATS_CLIENT_NAME" env-default:"nats-platform"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
