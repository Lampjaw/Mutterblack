package translatorplugin

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lampjaw/discordgobot"
)

type translatorPlugin struct {
	discordgobot.Plugin
}

func New() discordgobot.IPlugin {
	return &translatorPlugin{}
}

func (p *translatorPlugin) Commands() []discordgobot.CommandDefinition {
	return []discordgobot.CommandDefinition{
		discordgobot.CommandDefinition{
			CommandID:    "twanswate",
			Triggers: []string{
				"twanswate",
			},
			Description: "Twanswate the previous comment",
			Callback:    p.runTwanswateCommand,
		},
	}
}

func (p *translatorPlugin) Name() string {
	return "translator"
}

func (p *translatorPlugin) runTwanswateCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, message discordgobot.Message, args map[string]string, trigger string) {
	previousMessages, err := client.GetMessages(message.Channel(), 1, message.MessageID())

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	if previousMessages == nil || len(previousMessages) == 0 {
		p.RLock()
		client.SendMessage(message.Channel(), "Unable to find a message to translate.")
		p.RUnlock()
		return
	}

	var previousMessage = previousMessages[0]

	if client.IsMe(previousMessage) {
		return
	}

	replacer := strings.NewReplacer(
		"r", "w",
		"R", "W",
		"l", "w",
		"L", "W")

	translatedText := replacer.Replace(previousMessage.Message())

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	channel, err := client.Channel(message.Channel())
	guild, err := client.Guild(channel.GuildID)

	timestamp, err := previousMessage.Timestamp()

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    previousMessage.UserName(),
			IconURL: previousMessage.UserAvatar(),
		},
		Color:       0x070707,
		Description: translatedText,
		Timestamp:   timestamp.UTC().Format("2006-01-02T15:04:05-0700"),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("in #%s at %s", channel.Name, guild.Name),
		},
	}

	p.RLock()
	client.SendEmbedMessage(message.Channel(), embed)
	p.RUnlock()
}