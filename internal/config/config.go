package config

import "os"

// BotToken is the bot token for the bot
var BotToken = os.Getenv("TOKEN")
