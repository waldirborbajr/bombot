package handlers

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/botlog"
	"github.com/waldirborbajr/bombot/internal/constants"
)

func startGroup() string {
	return fmt.Sprintf(
		"Hey there! My name is ***BomBot*** - I'm here to help" +
			" you manage your groups! Hit /help to find out," +
			" more about how tp use me to my full potential. ",
	)
}

// HandlerGroup
func HandlerGroup(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := botlog.BotLog()

	logger.Info().Msg("HandlerFMS")

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
