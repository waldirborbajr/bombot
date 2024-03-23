package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/config"
	"github.com/waldirborbajr/bombot/internal/database"

	openai "github.com/sashabaranov/go-openai"
)

var (
	db   *database.Database
	err  error
	help string
)

var (
	chatMode      map[int64]string = make(map[int64]string)
	translateLang map[int64]string = make(map[int64]string)
	openaiClient  *openai.Client
)

func main() {
	// initialize log config
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err = database.New()
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	helpAux, err := os.ReadFile("help.md")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error reading help.md: %v", err)
		}
		log.Println("help.md not found")
	}
	help = string(helpAux)

	// --------------------------------
	/// Bot Initialization
	// --------------------------------

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	// OpenAI
	openaiClient = openai.NewClient(os.Getenv("OPENAI_KEY"))

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithDebug(),
	}

	telegramBotToken := config.BotToken
	if telegramBotToken == "" {
		log.Println("TELEGRAM_BOT_TOKEN environment variable is not set")
		return
	}

	b, err := bot.New(telegramBotToken, opts...)
	if nil != err {
		log.Fatalf("Error creating bot: %v", err)
	}

	webHookUrl := config.BotUrl
	if webHookUrl == "" {
		log.Println("webHook URL environment variable is not set")
		return
	}

	_, err = b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: webHookUrl,
	})
	if err != nil {
		log.Printf("Error on SetWebhook: %v", err)
		return
	}

	// Use StartWebhook instead of Start

	go b.StartWebhook(ctx)

	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "start",
				Description: "Explain the following text",
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

	go func() {
		err = http.ListenAndServe(":2000", b.WebhookHandler())
		if err != nil {
			log.Fatalf("Error Listening server: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("BomBot started")

	// call methods.DeleteWebhook if needed
	defer func() {
		_, err = b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
		if err != nil {
			log.Printf("Error on DeleteWebhook: %v", err)
			return
		}
	}()
}

// handler is a default handler that simply sends a message to the chat.
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.ChannelPost == nil {
		return
	}

	// Block to check for command
	switch {
	case update.ChannelPost.Text == "/id":
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.ChannelPost.Chat.ID,
			Text:   fmt.Sprintf("%d", update.ChannelPost.Chat.ID),
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
