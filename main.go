package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer func() {
		cancel()
	}()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	bot, err := bot.New(os.Getenv("TOKEN"), opts...)
	if err != nil {
		log.Fatalf("Error: recovering Token, %v", err)
		os.Exit(1)
	}

	go bot.StartWebhook(ctx)

	if err := http.ListenAndServe(":2020", bot.WebhookHandler()); err != nil {
		log.Fatalf("Error creating webhoob, %v", err)
		os.Exit(1)
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
