package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/botlog"
)

func start(update *models.Update) string {
	return fmt.Sprintf(
		"Hi ***%s***,\n"+
			"I am ***BomBot***, a Telegram Bot.\n\n"+
			"Send /help for getting info on how on use me!\n"+
			"Also you can send /source to get my source code to know how i'm built ;) and make sure to give a star to it; that makes my Devs to work more on O.S. projects like me :)\n\n"+
			"Hope you liked it !\n"+
			"Brought to You with ❤️ By @WaldirBorbaJr\n"+
			"Head towards @WaldirBorbaJr for any queries!", update.Message.From.Username)
}

func HandlerFMS(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger := botlog.BotLog()

	logger.Info().Msg("HandlerFMS")

	chatMode := make(map[int64]string)

	if update.Message == nil ||
		update.Message.Text == "" ||
		update.Message.From == nil {
		return
	}

	// Block to check for command
	switch {
	case len(update.Message.Entities) > 0:
		switch update.Message.Entities[0].Type == "bot_command" {
		case strings.HasPrefix(update.Message.Text, "/start"):
			b.SendMessage(ctx, &bot.SendMessageParams{
				ParseMode: "Markdown",
				ChatID:    update.Message.Chat.ID,
				Text:      start(update),
			})
			chatMode[update.Message.Chat.ID] = "start"
			return
		case strings.HasPrefix(update.Message.Text, "/explain"):
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "What do you want me to explain?",
			})
			chatMode[update.Message.Chat.ID] = "explain"
			return
		case strings.HasPrefix(update.Message.Text, "/help"):
			b.SendMessage(ctx, &bot.SendMessageParams{
				ParseMode: "Markdown",
				ChatID:    update.Message.Chat.ID,
				Text:      start(update),
			})
			return
		case strings.HasPrefix(update.Message.Text, "/image"):
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "What do you want me to generate?",
			})
			chatMode[update.Message.Chat.ID] = "image"
			return
		default:
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Unknown command. Use /help to get help.",
			})
			return
		}
	}

	// Block to check for the Finite State Machine
	switch chatMode[update.Message.Chat.ID] {
	case "start":
		number, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			logger.Error().Msgf("Error converting to number: %v", err)
			return
		}

		if number > 5 {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "I told you to type a number bertweem 1..5",
			})
			return
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Now type a number between 2..4",
			})
			chatMode[update.Message.Chat.ID] = "level1"
			return
		}
	case "level1":
		number, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			log.Println("Error converting to number")
			return
		}

		if number > 5 {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "I told you to type a number bertweem 2..4",
			})
			return
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Now type the number 3",
			})
			chatMode[update.Message.Chat.ID] = "level2"
			return
		}
	case "level2":
		number, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			logger.Error().Msgf("Error converting to number: %v", err)
			return
		}

		if number == 3 {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🎉 You Win 🏆",
			})
			chatMode[update.Message.Chat.ID] = "fini"
			return
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Now type the number 3",
			})
			chatMode[update.Message.Chat.ID] = "level2"
			return
		}
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
