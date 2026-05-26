package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/CodeHunt7/YT-AI-Bot/internal/config"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config load failed", "error", err)
		os.Exit(1)
	}

	logger.Info("bot starting", "env", cfg.Env)

	fmt.Println("Telegram bot token loaded")
}
