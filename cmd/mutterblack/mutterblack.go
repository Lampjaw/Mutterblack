package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/lampjaw/mutterblack/pkg/bot"
	"github.com/lampjaw/mutterblack/internal/pkg/plugin/invite"
	"github.com/lampjaw/mutterblack/internal/pkg/plugin/planetsidetwo"
	"github.com/lampjaw/mutterblack/internal/pkg/plugin/stats"
	"github.com/lampjaw/mutterblack/internal/pkg/plugin/translator"
)

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
	q := make(chan bool)

	if token == "" {
		fmt.Println("No token provided.")
		return
	}

	bot := bot.NewBot(token, clientID, ownerUserID)

	bot.RegisterPlugin(inviteplugin.New())
	bot.RegisterPlugin(statsplugin.New())
	bot.RegisterPlugin(planetsidetwoplugin.New())
	bot.RegisterPlugin(translatorplugin.New())

	bot.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	t := time.Tick(1 * time.Minute)

out:
	for {
		select {
		case <-q:
			break out
		case <-c:
			break out
		case <-t:
			bot.Save()
		}
	}

	bot.Save()
}
