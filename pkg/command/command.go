package command

import (
	"fmt"

	"github.com/lampjaw/mutterblack/pkg/discord"
)

func CommandHelp(client *discord.Discord, command string, arguments string, help string) string {
	if arguments != "" {
		return fmt.Sprintf("`%s%s %s` - %s", client.CommandPrefix(), command, arguments, help)
	}
	return fmt.Sprintf("`%s%s` - %s", client.CommandPrefix(), command, help)
}