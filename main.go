package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/config"
	"github.com/waldirborbajr/bombot/internal/database"
)

var (
	db   *database.Database
	err  error
	help string
)

// handler is a default handler that simply sends a message to the chat.
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := ""
	r := ""

	if update.Message == nil ||
		update.Message.Text == "" ||
		update.Message.From == nil {
		return
	}

	chatID := update.Message.Chat.ID
	message = update.Message.Text

	err := db.AddMessage("query", message, r)
	if err != nil {
		log.Printf("Error adding message to database: %v", err)
	}

	if strings.Contains(message[0:1], "/") {
		switch update.Message.Text {
		case "/start":
			message = "Welcome to the BomBot!"
		}
	}

	fmt.Println("ID: ", chatID)

	if !update.Message.From.IsBot {
		log.Println("From Bot: ")
	}

	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: chatID,
	// 	Text:   message,
	// })
}

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

	http.ListenAndServe(":2000", b.WebhookHandler())

	log.Println("BomBot started")

	// call methods.DeleteWebhook if needed
	_, err = b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
	if err != nil {
		log.Printf("Error on DeleteWebhook: %v", err)
		return
	}
}
