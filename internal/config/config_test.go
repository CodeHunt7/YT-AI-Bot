package config

import "testing"

func TestLoadRequiresTelegramBotToken(t *testing.T) {
	t.Setenv("APP_ENV", "local")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("DATABASE_URL", "postgres://ytbot:ytbot@localhost:5432/yt_ai_bot?sslmode=disable")

	_, err := Load()
	if err == nil {
		t.Fatalf("expected error for missing TELEGRAM_BOT_TOKEN")
	}
	if err.Error() != "TELEGRAM_BOT_TOKEN is required" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLoadRequiresDatabaseURL(t *testing.T) {
	t.Setenv("APP_ENV", "local")
	t.Setenv("TELEGRAM_BOT_TOKEN", "token")
	t.Setenv("DATABASE_URL", "")

	_, err := Load()
	if err == nil {
		t.Fatalf("expected error for missing DATABASE_URL")
	}
	if err.Error() != "DATABASE_URL is required" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLoadSuccess(t *testing.T) {
	t.Setenv("APP_ENV", "local")
	t.Setenv("TELEGRAM_BOT_TOKEN", "token")
	t.Setenv("DATABASE_URL", "postgres://ytbot:ytbot@localhost:5432/yt_ai_bot?sslmode=disable")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "local" {
		t.Fatalf("expected env local, got %q", cfg.Env)
	}
	if cfg.TelegramBotToken != "token" {
		t.Fatalf("expected telegram token to be loaded")
	}
	if cfg.DatabaseURL == "" {
		t.Fatalf("expected database url to be loaded")
	}
}
