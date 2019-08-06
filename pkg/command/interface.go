package command

import (
	"fmt"

	"github.com/lampjaw/mutterblack/pkg/discord"
)

type CommandDefinition struct {
	Description  string
	CommandID    string
	CommandGroup string
	Triggers     []string
	Arguments    []CommandDefinitionArgument
	Callback     func(client *discord.Discord, message discord.Message, args map[string]string, trigger string)
}

type CommandDefinitionArgument struct {
	Optional bool
	Pattern  string
	Alias    string
}

func (c *CommandDefinition) Help(client *discord.Discord) string {
	commandString := fmt.Sprintf("%s%s", client.CommandPrefix(), c.Triggers[0])

	if len(c.Arguments) > 0 {
		for _, argument := range c.Arguments {
			commandString = fmt.Sprintf("%s <%s>", commandString, argument.Alias)
		}
	}

	return fmt.Sprintf("`%s` - %s", commandString, c.Description)
}
