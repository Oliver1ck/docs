package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SERVER_HOST string `required:"true" envconfig:"SERVER_HOST"`
	SERVER_PORT string `required:"true" envconfig:"SERVER_PORT"`

	Postgres Postgres

	Logger Logger
}

func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
