package inviteplugin

import (
	"fmt"
	"strings"

	"github.com/lampjaw/mutterblack/internal/pkg/plugin"
	"github.com/lampjaw/mutterblack/pkg/command"
	"github.com/lampjaw/mutterblack/pkg/discord"
)

type invitePlugin struct {
	plugin.Plugin
}

func New() plugin.IPlugin {
	return &invitePlugin{}
}

func (p *invitePlugin) Commands() []command.CommandDefinition {
	return []command.CommandDefinition{
		command.CommandDefinition{
			CommandGroup: p.Name(),
			CommandID:    "invite",
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

func (p *invitePlugin) runInviteCommand(client *discord.Discord, message discord.Message, args map[string]string, trigger string) {
	if client.ApplicationClientID != "" {
		client.SendMessage(message.Channel(), fmt.Sprintf("Please visit <https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot> to add %s to your server.", client.ApplicationClientID, client.UserName()))
		return
	}
}

func discordInviteID(id string) string {
	id = strings.Replace(id, "://discordapp.com/invite/", "://discord.gg/", -1)
	id = strings.Replace(id, "https://discord.gg/", "", -1)
	id = strings.Replace(id, "http://discord.gg/", "", -1)
	return id
}