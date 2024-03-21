package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/waldirborbajr/bombot/internal/config"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(config.BotToken, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: "/https://93fb-2804-d55-8603-c700-8418-afa0-2687-d854.ngrok-free.app/webhook",
	})

	go func() {
		http.ListenAndServe(":9090", b.WebhookHandler())
	}()

	// Use StartWebhook instead of Start
	b.StartWebhook(ctx)

	log.Println("BomBot started...")

	// call methods.DeleteWebhook if needed
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil ||
		update.Message.Text == "" ||
		update.Message.From == nil {
		return
	}

	log.Printf("From: %v\nText: %v\n", update.Message.From, update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
