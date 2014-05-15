package config

import (
	"github.com/kelseyhightower/envconfig"
)

// LoadEnv loads all configuration from the environment
func LoadEnv() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("newrelic", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
