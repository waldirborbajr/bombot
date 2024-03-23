package handlers

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

func helpGroup() string {
	return fmt.Sprint(
		"Add me to your group with the following permissions and I'll handle the rest!",
		"\n - Ban Permissions: To ban the channels",
		"\n - Delete Message Permissions: To delete the messages sent by channel",

		"\n\n***Some Tips:***",
		"\n1. To ignore a channel use /ignore by replying a message from that channel or you can pass a channel id. for more help type /ignore.",
		"\n2. To unignore a channel use /unignore by replying a message from that channel or you can pass a channel id. for more help type /unignore.",
		"\n3. To get the list of all ignored channel use /ignorelist.",

		"\n\n***Available Commands:***",
		"\n/start - ✨ display start message.",
		"\n/ignore - ✅ unban and allow that user to sending message as channel (admin only).",
		"\n/ignorelist - 📋 get list ignored channel.",
		"\n/unignore - ⛔️ ban an unallow that user to sending message as channel (admin only).",
		"\n/source - 📚 get source code.",
	)
}

// HandlerGroup a default handler that simply sends a message to the chat.
func HandlerGroup(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info().Msg("HandlerGroup")

	if update.ChannelPost == nil {
		return
	}

	// Block to check for command
	switch update.ChannelPost.Text {
	case "/id":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.ChannelPost.Chat.ID,
			Text:   fmt.Sprintf("This channel id is: %d", update.ChannelPost.Chat.ID),
		})
		return
	case "/help":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ParseMode: "Markdown",
			ChatID:    update.ChannelPost.Chat.ID,
			Text:      helpGroup(),
		})
		return
	}
}
