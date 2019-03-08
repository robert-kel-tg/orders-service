package config

import "github.com/kelseyhightower/envconfig"

type (
	// Spec ...
	Spec struct {
		DatabaseDSN string `envconfig:"DATABASE_DSN"`
		Port        int    `envconfig:"PORT"`
	}
)

func LoadEnv() (*Spec, error) {
	var (
		config Spec
	)
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
