package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/bombot/internal/constants"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// botName := update.Message.Chat.Username
	botName := "NoCoderaBot"

	log.Info().Msg(botName)

	keyb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text: "Add me to your chat!",
					URL:  "https://t.me/" + botName + "?startgroup=None",
				},
			},
		},
	}
	// keyb := &models.InlineKeyboardMarkup{
	// 	InlineKeyboard: [][]models.InlineKeyboardButton{
	// 		{
	// 			{Text: "Add me to your chat!", CallbackData: "button add"},
	// 		},
	// 	},
	// }

	// switch update.CallbackQuery.Data {
	// case "add to chat":
	// 	log.Info().Msg("add to chat")
	// }

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
