package fms

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// MenuCommands sets the menu commands for the bot
func MenuCommands(ctx context.Context, b *bot.Bot) {
	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "start",
				Description: "Explain the following text",
			},
			{
				Command:     "joke",
				Description: "Tell me a joke",
			},
			{
				Command:     "help",
				Description: "Cry for help",
			},
			{
				Command:     "explain",
				Description: "Explain the following text",
			},
			{
				Command:     "translate",
				Description: "Translate from any language to any other language",
			},
			{
				Command:     "image",
				Description: "Generate an image from text",
			},
		},
	})
}
