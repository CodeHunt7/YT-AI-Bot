package config

import "os"

type Config struct {
	Env string
}

func Load() Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	return Config{
		Env: env,
	}
}
