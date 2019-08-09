package weatherplugin

//
// THIS IS DEPRECATED AND WILL BE REMOVED IN THE FUTURE
//

import "github.com/lampjaw/discordgobot"

type weatherPlugin struct {
	discordgobot.Plugin
}

func New() discordgobot.IPlugin {
	return &weatherPlugin{}
}

func (p *weatherPlugin) Commands() []discordgobot.CommandDefinition {
	return []discordgobot.CommandDefinition{
		discordgobot.CommandDefinition{
			CommandID: "weather",
			Triggers: []string{
				"w",
				"wf",
			},
			Description: "Discontinued",
			Callback:    p.runCommand,
		},
	}
}

func (p *weatherPlugin) Name() string {
	return "Weather"
}

func (p *weatherPlugin) runCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, message discordgobot.Message, args map[string]string, trigger string) {
	p.Lock()
	client.SendMessage(message.Channel(), "Weather has been removed from Mutterblack. If this is a feature you would like to continue to use on your server, please check out my new bot Weatherman! <https://github.com/Lampjaw/Weatherman>")
	p.Unlock()
}
