package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

type Command struct {
	Handler InteractionCreateEventHandler
	Name    string
}

var (
	Commands []*Command
)

func registerCommand(command *Command) {
	Commands = append(Commands, command)
}

func linkCommands(s *discordgo.Session) {
	for _, command := range Commands {
		color.Magenta("[Commands] Registering %s", command)

		handler := InjectCommandMiddleware(command)

		color.Red("[Commands] %s Interface type: %T", command.Name, handler)

		s.AddHandler(handler)
	}
}
