package bot

import (
	"context"
	"log/slog"

	"github.com/CodeHunt7/YT-AI-Bot/internal/app"
	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	client     *tgbot.Bot
	logger     *slog.Logger
	onboarding *app.OnboardingService
}

func New(token string, logger *slog.Logger) (*Bot, error) {
	client, err := tgbot.New(token)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		client:     client,
		logger:     logger,
		onboarding: app.NewOnboardingService(),
	}

	client.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypeExact, b.handleStart)
	client.RegisterHandler(tgbot.HandlerTypeMessageText, "", tgbot.MatchTypePrefix, b.handleText)

	return b, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.logger.Info("telegram bot started")
	b.client.Start(ctx)
}

func (b *Bot) handleStart(ctx context.Context, client *tgbot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}

	text := b.onboarding.HandleStart(update.Message.From.ID)
	b.sendMessage(ctx, client, update.Message.Chat.ID, text)
}

func (b *Bot) handleText(ctx context.Context, client *tgbot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}
	if update.Message.Text == "/start" {
		return
	}

	text := b.onboarding.HandleText(update.Message.From.ID, update.Message.Text)
	b.sendMessage(ctx, client, update.Message.Chat.ID, text)
}

func (b *Bot) sendMessage(ctx context.Context, client *tgbot.Bot, chatID int64, text string) {
	_, err := client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		b.logger.Error("send message failed", "error", err)
	}
}
