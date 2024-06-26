package constants

import "github.com/go-telegram/bot"

// This file contains all the text messages used. Only reason for having this in
// separate file is these can't be declared as consts and looks really awful.
// I found this workaround as passing the complete as string doesn't work,
// I don't know if that was intentional or not. Gotta find a cleaner way to do this.

// Default message for any unrecognized command.
var DefaultMessage = "Sorry, I didn't understand that command. " +
	"Please make sure you've entered a valid instruction, or type /help for assistance."

// Messages used with start command.
var StartMessage = bot.EscapeMarkdown(
	"Hello! My name is BomBot - I'm here to help you manage " +
		"your groups! Hit /help to find out more about how " +
		"to use me to my full potential.\n\n" +
		"Remember type at any moment /help for more info.",
)

// Messages used with help command.
var HelpMessage = "*Note:* " +
	bot.EscapeMarkdown(
		"The bot will fetch some of the recent torrents, so be specific with search query.\n\n",
	) +
	"*Available commands:*" +
	bot.EscapeMarkdown(
		"\n- /start - To check whether the bot is alive.\n- /help - To display this message.\n- /magnet - To get torrent info from ",
	) +
	"*Nyaa* and *Sukebei* " +
	bot.EscapeMarkdown("using ID.") +
	bot.EscapeMarkdown("\n- /nyaa - For searching torrents on Nyaa.") +
	bot.EscapeMarkdown("\n- /sukebei - For searching torrents on Sukebei.")

// Messages used with magnet command.
var (
	MagnetMessage    = "Where do you wanna search?"
	NoIDMessage      = "No ID provided."
	InvalidIDMessage = "Invalid ID!"
)

// Messages used with nyaa and sukebei search commands.
var (
	MissingQueryMessage = "Missing search query!"
	CatMessage          = "Please choose one of the following categories: "
	SubCatMessage       = "Please choose one of the following sub-categories: "
	NoResultMessage     = "No results found! It's like they pulled a disappearing act. Try another search and let the hunt continue!"
)

// Common messages.
var SomethingWentWrong = "Something went wrong!"
