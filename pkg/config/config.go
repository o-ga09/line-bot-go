package config

import "github.com/caarlos0/env"

type Config struct {
	Env              string `env:"ENV" envDefault:"dev"`
	Port             string `env:"PORT" envDefault:"80"`
	Database_url     string `env:"DATABASE_URL" envDefult:""`
	ProjectID        string `env:"PROJECTID" envDefault:""`
	LineChannelscret string `env:"LINE_CHANNEL_SECRET"`
	LineAccesstoken  string `env:"LINE_ACCESS"`
	LineUserId       string `env:"LINE_USER_ID"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
