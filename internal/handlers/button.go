package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

func CallbackButtonMenuHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info().Msg("CallbackMenuHandler called (Button)")

	switch update.CallbackQuery.Data {
	case "button add":
		log.Info().Msg("Add to Channel or Group")
	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "unknown button",
		})
	}
}
