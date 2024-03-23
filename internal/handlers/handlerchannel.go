package handlers

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/botlog"
)

func help() string {
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

// handler is a default handler that simply sends a message to the chat.
func HandlerChannel(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := botlog.BotLog()

	logger.Info().Msg("HandlerChannel")

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
	case "/ignore":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ParseMode: "Markdown",
			ChatID:    update.ChannelPost.Chat.ID,
			Text:      help(),
		})
		return
	case "/help":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ParseMode: "Markdown",
			ChatID:    update.ChannelPost.Chat.ID,
			Text:      help(),
		})
		return
	}

	// msg, _ := json.Marshal(update)
	// log.Default().Println(string(msg))
	//
	// if len(update.Message.Entities) > 0 {
	// 	if update.Message.Entities[0].Type == "bot_command" &&
	// 		strings.HasPrefix(update.Message.Text, "/explain") {
	// 		b.SendMessage(ctx, &bot.SendMessageParams{
	// 			ChatID: update.Message.Chat.ID,
	// 			Text:   "What do you want me to explain?",
	// 		})
	// 		chatMode[update.Message.Chat.ID] = "explain"
	// 		return
	// 	}
	// 	if update.Message.Entities[0].Type == "bot_command" &&
	// 		strings.HasPrefix(update.Message.Text, "/translate") {
	// 		lang := ""
	// 		if len(
	// 			update.Message.Text,
	// 		) > update.Message.Entities[0].Offset+update.Message.Entities[0].Length {
	// 			lang = update.Message.Text[update.Message.Entities[0].Offset+update.Message.Entities[0].Length:]
	// 		} else {
	// 			b.SendMessage(ctx, &bot.SendMessageParams{
	// 				ChatID: update.Message.Chat.ID,
	// 				Text:   "Select the language you want me to translate to:",
	// 			})
	// 			chatMode[update.Message.Chat.ID] = "ask_language"
	// 			return
	// 		}
	// 		b.SendMessage(ctx, &bot.SendMessageParams{
	// 			ChatID: update.Message.Chat.ID,
	// 			Text:   "What do you want me to translate?",
	// 		})
	// 		chatMode[update.Message.Chat.ID] = "translate"
	// 		translateLang[update.Message.Chat.ID] = lang
	//
	// 		return
	// 	}
	//
	// 	if update.Message.Entities[0].Type == "bot_command" &&
	// 		strings.HasPrefix(update.Message.Text, "/image") {
	// 		b.SendMessage(ctx, &bot.SendMessageParams{
	// 			ChatID: update.Message.Chat.ID,
	// 			Text:   "What do you want me to generate?",
	// 		})
	// 		chatMode[update.Message.Chat.ID] = "image"
	// 		return
	// 	}
	//
	// }
	//
	// if chatMode[update.Message.Chat.ID] == "ask_language" {
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID: update.Message.Chat.ID,
	// 		Text:   "What do you want me to translate?",
	// 	})
	// 	chatMode[update.Message.Chat.ID] = "translate"
	// 	translateLang[update.Message.Chat.ID] = update.Message.Text
	// 	return
	// }
	//
	// if chatMode[update.Message.Chat.ID] == "translate" {
	// 	resp, err := openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
	// 		Model: "gpt-3.5-turbo",
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role: "user",
	// 				Content: fmt.Sprintf(
	// 					"Translate `%s` to `%s`",
	// 					update.Message.Text,
	// 					translateLang[update.Message.Chat.ID],
	// 				),
	// 			},
	// 		},
	// 	})
	// 	var msg string
	// 	if err != nil {
	// 		log.Default().Println("Error:", err)
	// 		msg = "I'm sorry, I couldn't translate that. Please try again."
	// 	} else {
	// 		msg = resp.Choices[0].Message.Content
	// 	}
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID: update.Message.Chat.ID,
	// 		Text:   msg,
	// 	})
	// 	delete(chatMode, update.Message.Chat.ID)
	// 	delete(translateLang, update.Message.Chat.ID)
	// 	return
	// }
	//
	// if chatMode[update.Message.Chat.ID] == "explain" {
	// 	resp, err := openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
	// 		Model: "gpt-3.5-turbo",
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    "user",
	// 				Content: "Please explain:\n" + update.Message.Text,
	// 			},
	// 		},
	// 	})
	// 	var msg string
	// 	if err != nil {
	// 		log.Default().Println("Error:", err)
	// 		msg = "I'm sorry, I couldn't explain that. Please try again."
	// 	} else {
	// 		msg = resp.Choices[0].Message.Content
	// 	}
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID: update.Message.Chat.ID,
	// 		Text:   msg,
	// 	})
	// 	delete(chatMode, update.Message.Chat.ID)
	// 	return
	// }
	//
	// if chatMode[update.Message.Chat.ID] == "image" {
	// 	resp, err := openaiClient.CreateImage(ctx, openai.ImageRequest{
	// 		Prompt: update.Message.Text,
	// 		N:      1,
	// 	})
	//
	// 	var msg string
	// 	if err != nil {
	// 		log.Default().Println("Error:", err)
	// 		msg = "I'm sorry, I couldn't explain that. Please try again."
	// 		b.SendMessage(ctx, &bot.SendMessageParams{
	// 			ChatID: update.Message.Chat.ID,
	// 			Text:   msg,
	// 		})
	// 	}
	//
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID: update.Message.Chat.ID,
	// 		Text:   resp.Data[0].URL,
	// 	})
	// 	delete(chatMode, update.Message.Chat.ID)
	// 	return
	// }

	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   update.Message.Text,
	// })
}
