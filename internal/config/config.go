package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	TelegramBotToken string
	DatabaseURL      string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return Config{}, errors.New("TELEGRAM_BOT_TOKEN is required")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}

	return Config{
		Env:              env,
		TelegramBotToken: token,
		DatabaseURL:      databaseURL,
	}, nil
}
