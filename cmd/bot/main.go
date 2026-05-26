package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/CodeHunt7/YT-AI-Bot/internal/config"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("bot starting", "env", cfg.Env)

	fmt.Println("YT AI Bot skeleton is running")
}
