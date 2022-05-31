package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"TODO_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"TODO_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"TODO_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"TODO_DB_USER" envDefault:"todo"`
	DBPassword string `env:"TODO_DB_PASSWORD" envDefault:"todo"`
	DBName     string `env:"TODO_DB_NAME" envDefault:"todo"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
