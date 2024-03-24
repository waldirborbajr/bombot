package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/bombot/internal/constants"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info().Msg("DefaultHandler called")

	if update.Message == nil ||
		update.Message.Text == "" ||
		update.Message.From == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   constants.DefaultMessage,
	})
}
