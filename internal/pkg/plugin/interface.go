package plugin

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/lampjaw/mutterblack/pkg/discord"
	"github.com/lampjaw/mutterblack/pkg/command"
)

type IPlugin interface {
	Name() string
	Load(*discord.Discord, []byte) error
	Save() ([]byte, error)
	Help(*discord.Discord, discord.Message, bool) []string
	Message(*discord.Discord, discord.Message) error
	Commands() []command.CommandDefinition
}

type Plugin struct {
	sync.RWMutex
}

func (p *Plugin) Commands() []command.CommandDefinition {
	return nil
}

func (p *Plugin) Name() string {
	return ""
}

func (p *Plugin) Load(client *discord.Discord, data []byte) error {
	if data != nil {
		if err := json.Unmarshal(data, p); err != nil {
			log.Println("Error loading data", err)
		}
	}

	return nil
}

func (p *Plugin) Save() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Plugin) Help(client *discord.Discord, message discord.Message, detailed bool) []string {
	return nil
}

func (p *Plugin) Message(client *discord.Discord, message discord.Message) error {
	return nil
}
