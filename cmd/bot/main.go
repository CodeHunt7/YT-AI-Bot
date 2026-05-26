package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	internalbot "github.com/CodeHunt7/YT-AI-Bot/internal/bot"
	"github.com/CodeHunt7/YT-AI-Bot/internal/config"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config load failed", "error", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	b, err := internalbot.New(cfg.TelegramBotToken, logger)
	if err != nil {
		logger.Error("telegram bot init failed", "error", err)
		os.Exit(1)
	}

	logger.Info("bot starting", "env", cfg.Env)
	b.Start(ctx)
}
