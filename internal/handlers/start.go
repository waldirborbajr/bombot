package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/constants"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	keyb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Add me to your chat!", URL: "data.Data.Link"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        constants.StartMessage,
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: keyb,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}
