package inviteplugin

import (
	"fmt"

	"github.com/lampjaw/discordgobot"
)

type invitePlugin struct {
	discordgobot.Plugin
}

func New() discordgobot.IPlugin {
	return &invitePlugin{}
}

func (p *invitePlugin) Commands() []*discordgobot.CommandDefinition {
	return []*discordgobot.CommandDefinition{
		&discordgobot.CommandDefinition{
			CommandID: "invite",
			Triggers: []string{
				"invite",
			},
			Description: "Get an invite link to add this bot to your server!",
			Callback:    p.runInviteCommand,
		},
	}
}

func (p *invitePlugin) Name() string {
	return "Invite"
}

func (p *invitePlugin) runInviteCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, payload discordgobot.CommandPayload) {
	message := payload.Message

	if bot.Config != nil && bot.Config.ClientID != "" {
		client.SendMessage(message.Channel(),
			fmt.Sprintf("Please visit <https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot> to add %s to your server.",
				bot.Config.ClientID, client.UserName()))
		return
	}
}
