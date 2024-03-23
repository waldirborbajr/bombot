package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/rs/zerolog"
	"github.com/waldirborbajr/bombot/internal/config"
	"github.com/waldirborbajr/bombot/internal/database"
	"github.com/waldirborbajr/bombot/internal/fms"

	openai "github.com/sashabaranov/go-openai"
)

var (
	db  *database.Database
	err error
	// help   string
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Logger()
)

var (
	// chatMode      map[int64]string = make(map[int64]string)
	translateLang map[int64]string = make(map[int64]string)
	openaiClient  *openai.Client
)

func main() {
	db, err = database.New()
	if err != nil {
		logger.Error().Msgf("Error creating database: %v", err)
	}

	// --------------------------------
	/// Bot Initialization
	// --------------------------------

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	// OpenAI
	openaiClient = openai.NewClient(os.Getenv("OPENAI_KEY"))

	opts := []bot.Option{
		bot.WithDefaultHandler(fms.Handler),
		bot.WithDebug(),
	}

	telegramBotToken := config.BotToken
	if telegramBotToken == "" {
		logger.Error().Msg("TELEGRAM_BOT_TOKEN environment variable is not set")
		return
	}

	b, err := bot.New(telegramBotToken, opts...)
	if nil != err {
		logger.Error().Msgf("Error creating bot: %v", err)
	}

	webHookUrl := config.BotUrl
	if webHookUrl == "" {
		logger.Error().Msgf("webHook URL environment variable is not set: %v", err)
		return
	}

	_, err = b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: webHookUrl,
	})
	if err != nil {
		logger.Error().Msgf("Error on SetWebhook: %v", err)
		return
	}

	// Use StartWebhook instead of Start

	go b.StartWebhook(ctx)

	fms.MenuCommands(ctx, b)

	go func() {
		err = http.ListenAndServe(":2000", b.WebhookHandler())
		if err != nil {
			logger.Error().Msgf("Error Listening server: %v", err)
		}
	}()

	// call methods.DeleteWebhook if needed
	defer func() {
		_, err = b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
		if err != nil {
			logger.Error().Msgf("Error on DeleteWebhook: %v", err)
			return
		}
	}()

	// <-ctx.Done()
	select {
	case <-ctx.Done():
		logger.Info().Msg("BomBot is shutting down...")
	}
}
