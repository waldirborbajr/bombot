package config

import "os"

// BotToken is the bot token for the bot
var (
	BotToken = os.Getenv("TOKEN")
	BotUrl   = os.Getenv("BOT_URL")
	BOT_FLAG = os.Getenv("BOT_FLAG")
)
