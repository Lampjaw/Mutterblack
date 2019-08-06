package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sort"

	"github.com/lampjaw/mutterblack/pkg/discord"
	"github.com/lampjaw/mutterblack/pkg/command"
	"github.com/lampjaw/mutterblack/internal/pkg/plugin"
)

type Bot struct {
	Client          *discord.Discord
	Plugins         map[string]plugin.IPlugin
	messageChannels []chan discord.Message
}

func (b *Bot) Open() {
	if messageChan, err := b.Client.Open(); err == nil {
		for _, plugin := range b.Plugins {
			plugin.Load(b.Client, b.getData(plugin))
		}
		go b.listen(messageChan)
	} else {
		log.Printf("Error creating discord service: %v\n", err)
	}
}

func (b *Bot) Save() {
	if err := os.Mkdir("data", os.ModePerm); err != nil {
		if !os.IsExist(err) {
			log.Println("Error creating service directory.")
		}
	}
	for _, plugin := range b.Plugins {
		if data, err := plugin.Save(); err != nil {
			log.Printf("Error saving plugin %s. %v", plugin.Name(), err)
		} else if data != nil {
			if err := ioutil.WriteFile("data/"+plugin.Name(), data, os.ModePerm); err != nil {
				log.Printf("Error saving plugin %s. %v", plugin.Name(), err)
			}
		}
	}
}

func (b *Bot) getData(plugin plugin.IPlugin) []byte {
	fileName := "data/" + plugin.Name()

	if b, err := ioutil.ReadFile(fileName); err == nil {
		return b
	}

	return nil
}

func (b *Bot) RegisterPlugin(plugin plugin.IPlugin) {
	if b.Plugins[plugin.Name()] != nil {
		log.Println("Plugin with that name already registered", plugin.Name())
	}
	b.Plugins[plugin.Name()] = plugin
}

func (b *Bot) listen(messageChan <-chan discord.Message) {
	log.Printf("Listening")
	for {
		message := <-messageChan

		if HandleCommandsRequest(b, message) {
			return
		}

		plugins := b.Plugins
		for _, plugin := range plugins {
			go plugin.Message(b.Client, message)
			if !b.Client.IsMe(message) {
				go findCommandMatch(b, plugin, message)
			}
		}
	}
}

func findCommandMatch(b *Bot, plugin plugin.IPlugin, message discord.Message) {
	if plugin.Commands() == nil || message.Message() == "" {
		return
	}

	for _, commandDefinition := range plugin.Commands() {
		for _, trigger := range commandDefinition.Triggers {
			var trig = b.Client.CommandPrefix() + trigger
			var parts = strings.Split(message.Message(), " ")

			if parts[0] == trig {
				log.Printf("<%s> %s: %s\n", message.Channel(), message.UserName(), message.Message())

				if commandDefinition.Arguments == nil {
					commandDefinition.Callback(b.Client, message, nil, trigger)
					return
				}

				parsedArgs := extractCommandArguments(message, trig, commandDefinition.Arguments)

				if parsedArgs != nil {
					commandDefinition.Callback(b.Client, message, parsedArgs, trigger)
					return
				}
			}
		}
	}
}

func extractCommandArguments(message discord.Message, trigger string, arguments []command.CommandDefinitionArgument) map[string]string {
	var argPatterns []string
	for _, argument := range arguments {
		argPatterns = append(argPatterns, fmt.Sprintf("(?P<%s>%s)", argument.Alias, argument.Pattern))
	}
	var pattern = fmt.Sprintf("^%s$", strings.Join(argPatterns, " "))

	var trimmedContent = strings.TrimPrefix(message.Message(), fmt.Sprintf("%s ", trigger))
	pat := regexp.MustCompile(pattern)
	argsMatch := pat.FindStringSubmatch(trimmedContent)

	parsedArgs := make(map[string]string)

	if argsMatch == nil || len(argsMatch) == 1 {
		return nil
	}

	for i := 1; i < len(argsMatch); i++ {
		parsedArgs[pat.SubexpNames()[i]] = argsMatch[i]
	}

	if len(parsedArgs) != len(arguments) {
		return nil
	}

	return parsedArgs
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func HandleCommandsRequest(b *Bot, message discord.Message) bool {
	var trig = b.Client.CommandPrefix() + "commands"
	var parts = strings.Split(message.Message(), " ")

	if parts[0] != trig {
		return false
	}

	help := []string{}

	for _, plugin := range b.Plugins {
		var h []string

		if plugin.Commands() == nil {
			h = plugin.Help(b.Client, message, false)
		} else {
			for _, commandDefinition := range plugin.Commands() {
				h = append(h, commandDefinition.Help(b.Client))
			}
		}

		if h != nil && len(h) > 0 {
			help = append(help, h...)
		}
	}

	sort.Strings(help)

	if len(help) == 0 {
		help = []string{fmt.Sprintf("Unknown topic: %s", parts[0])}
	}

	b.Client.SendMessage(message.Channel(), strings.Join(help, "\n"))

	return true
}