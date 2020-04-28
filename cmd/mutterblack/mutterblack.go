package main

import (
	"fmt"
	"os"
	"os/signal"

	commandplugin "mutterblack/pkg/plugins/command"
	inviteplugin "mutterblack/pkg/plugins/invite"
	planetsidetwoplugin "mutterblack/pkg/plugins/planetsidetwo"
	statsplugin "mutterblack/pkg/plugins/stats"
	translatorplugin "mutterblack/pkg/plugins/translator"

	"github.com/lampjaw/discordgobot"
)

// VERSION of Mutterblack
const VERSION = "3.0.0"

func init() {
	token = os.Getenv("Token")
	clientID = os.Getenv("ClientId")
	ownerUserID = os.Getenv("OwnerUserId")
}

var token string
var clientID string
var ownerUserID string
var buffer = make([][]byte, 0)

func main() {
	if token == "" {
		fmt.Println("No token provided.")
		return
	}

	commandPlugin := commandplugin.New()

	config := &discordgobot.GobotConf{
		OwnerUserID: ownerUserID,
		ClientID:    clientID,
		CommandPrefixFunc: func(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, message discordgobot.Message) string {
			channel, _ := client.Channel(message.Channel())
			prefix, err := commandPlugin.GetGuildPrefix(channel.GuildID)

			if err != nil || prefix == nil {
				return "?"
			}

			return *prefix
		},
	}

	bot, err := discordgobot.NewBot(token, config, nil)

	if err != nil {
		fmt.Sprintln("Unable to create bot: %s", err)
		return
	}

	bot.RegisterPlugin(commandPlugin)
	bot.RegisterPlugin(inviteplugin.New())
	bot.RegisterPlugin(statsplugin.New(VERSION))
	bot.RegisterPlugin(planetsidetwoplugin.New())
	bot.RegisterPlugin(translatorplugin.New())

	bot.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

out:
	for {
		select {
		case <-c:
			break out
		}
	}
}
