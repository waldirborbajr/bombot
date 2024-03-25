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

		chatID := update.CallbackQuery.Message.Message.Chat.ID

		totalMembers, err := b.GetChatMemberCount(
			ctx,
			&bot.GetChatMemberCountParams{ChatID: chatID},
		)
		if err != nil {
			log.Err(err).Msg("Error")
		}

		log.Info().Msgf("Total Members: %d", totalMembers)

		// admins, err := b.GetChatAdministrators(
		// 	ctx,
		// 	&bot.GetChatAdministratorsParams{ChatID: chatID},
		// )
		// if err != nil {
		// 	log.Err(err).Msg("Error")
		// }
		//
		// log.Info().Msg(admins.ChatMember.Owner)

		return

	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "unknown button",
		})
	}
}
