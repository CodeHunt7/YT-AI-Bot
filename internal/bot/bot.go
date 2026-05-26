package bot

import (
	"context"
	"log/slog"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	client *tgbot.Bot
	logger *slog.Logger
}

func New(token string, logger *slog.Logger) (*Bot, error) {
	client, err := tgbot.New(token)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		client: client,
		logger: logger,
	}

	client.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypeExact, b.handleStart)

	return b, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.logger.Info("telegram bot started")
	b.client.Start(ctx)
}

func (b *Bot) handleStart(ctx context.Context, client *tgbot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	_, err := client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет. Я ИИ-продюсер для YouTube-авторов. Пока умею только запускаться.",
	})
	if err != nil {
		b.logger.Error("send start message failed", "error", err)
	}
}
