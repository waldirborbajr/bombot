package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/bombot/internal/config"
	"github.com/waldirborbajr/bombot/internal/database"
	"github.com/waldirborbajr/bombot/internal/handlers"
	"github.com/waldirborbajr/bombot/internal/menu"
)

var (
	db       *database.Database
	err      error
	BOT_FLAG string
)

// chatMode      map[int64]string = make(map[int64]string)
var translateLang map[int64]string = make(map[int64]string)

func main() {
	BOT_FLAG = config.BOT_FLAG

	log.Info().Msg(BOT_FLAG)

	db, err = database.New()
	if err != nil {
		log.Error().Msgf("Error creating database: %v", err)
	}

	// --------------------------------
	/// Bot Initialization
	// --------------------------------

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	opts := []bot.Option{}

	switch BOT_FLAG {
	case "fms":
		opts = []bot.Option{
			bot.WithDefaultHandler(handlers.HandlerFMS),
			bot.WithDebug(),
		}
	case "channel":
		opts = []bot.Option{
			bot.WithDefaultHandler(handlers.HandlerChannel),
			bot.WithDebug(),
		}
	case "group":
		opts = []bot.Option{
			bot.WithDefaultHandler(handlers.DefaultHandler),

			bot.WithCallbackQueryDataHandler(
				"button",
				bot.MatchTypePrefix,
				handlers.CallbackButtonMenuHandler,
			),

			bot.WithCallbackQueryDataHandler(
				"magnet",
				bot.MatchTypePrefix,
				handlers.MagnetCallbackHandler,
			),
			bot.WithCallbackQueryDataHandler(
				"nyaa",
				bot.MatchTypePrefix,
				handlers.SearchCallbackHandler,
			),
			bot.WithCallbackQueryDataHandler(
				"sukebei",
				bot.MatchTypePrefix,
				handlers.SearchCallbackHandler,
			),
			// bot.WithCallbackQueryDataHandler("addToGroup", bot.MatchType, handler bot.HandlerFunc)
			bot.WithDebug(),
		}
	}

	telegramBotToken := config.BotToken
	if telegramBotToken == "" {
		log.Fatal().Msg("TELEGRAM_BOT_TOKEN environment variable is not set")
		return
	}

	b, err := bot.New(telegramBotToken, opts...)
	if nil != err {
		log.Error().Msgf("Error creating bot: %v", err)
	}

	webHookUrl := config.BotUrl
	if webHookUrl == "" {
		log.Fatal().Msg("webHook URL environment variable is not set")
		return
	}

	_, err = b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: webHookUrl,
	})
	if err != nil {
		log.Fatal().Msgf("Error on SetWebhook: %v", err)
		return
	}

	// Use StartWebhook instead of Start

	go b.StartWebhook(ctx)

	switch BOT_FLAG {
	case "fms":
		menu.MenuCommandsFMS(ctx, b)
	case "channel":
		menu.MenuCommandsChannel(ctx, b)
	case "group":
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			"/start",
			bot.MatchTypeExact,
			handlers.StartHandler,
		)
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			"/help",
			bot.MatchTypeExact,
			handlers.HelpHandler,
		)
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			"/magnet",
			bot.MatchTypePrefix,
			handlers.MagnetHandler,
		)
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			"/nyaa",
			bot.MatchTypePrefix,
			handlers.SearchHandler,
		)
		b.RegisterHandler(
			bot.HandlerTypeMessageText,
			"/sukebei",
			bot.MatchTypePrefix,
			handlers.SearchHandler,
		)
	}

	go func() {
		err = http.ListenAndServe(":2000", b.WebhookHandler())
		if err != nil {
			log.Fatal().Msgf("Error Listening server: %v", err)
		}
	}()

	// call methods.DeleteWebhook if needed
	defer func() {
		_, err = b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
		if err != nil {
			log.Error().Msgf("Error on DeleteWebhook: %v", err)
			return
		}
	}()

	// <-ctx.Done()
	select {
	case <-ctx.Done():
		log.Info().Msg("BomBot is shutting down...")
	}
}
