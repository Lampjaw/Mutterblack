package bot

import (
	"fmt"

	"github.com/lampjaw/mutterblack/internal/pkg/plugin"
	"github.com/lampjaw/mutterblack/pkg/discord"
)

func NewBot(token string, clientId string, ownerUserId string) *Bot {
	if token == "" {
		fmt.Println("No token provided. Please run: mutterblack -t <bot token>")
		return nil
	}

	bot := &Bot{
		Plugins: make(map[string]plugin.IPlugin, 0),
		Client:  discord.NewDiscord("Bot " + token),
	}

	bot.Client.ApplicationClientID = clientId
	bot.Client.OwnerUserID = ownerUserId

	return bot
}
