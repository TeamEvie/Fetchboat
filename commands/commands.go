package commands

import "github.com/bwmarrin/discordgo"

func Inject(s *discordgo.Session) {
	registerCommand(PingCommand)

	linkCommands(s)
}
