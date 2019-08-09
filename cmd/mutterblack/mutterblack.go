package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/lampjaw/discordgobot"
	inviteplugin "github.com/lampjaw/mutterblack/pkg/plugins/invite"
	planetsidetwoplugin "github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo"
	statsplugin "github.com/lampjaw/mutterblack/pkg/plugins/stats"
	translatorplugin "github.com/lampjaw/mutterblack/pkg/plugins/translator"
)

// VERSION of Mutterblack
const VERSION = "2.0.0"

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

	bot, err := discordgobot.NewBot(token, "?", clientID, ownerUserID)

	if err != nil {
		fmt.Sprintln("Unable to create bot: %s", err)
		return
	}

	bot.RegisterPlugin(inviteplugin.New())
	bot.RegisterPlugin(statsplugin.New())
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
