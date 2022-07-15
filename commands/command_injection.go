package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

type InteractionCreateEventHandler func(*discordgo.Session, *discordgo.InteractionCreate)

func InjectCommandMiddleware(cmd *Command) InteractionCreateEventHandler {
	color.Magenta("[Commands] Injecting handler for %s", cmd.Name)
	return func(s *discordgo.Session, m *discordgo.InteractionCreate) {
		if m.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if m.ApplicationCommandData().Name == cmd.Name {
			color.Cyan("[Commands] Executing %s | 🧍 %s | ⏰ %s", cmd.Name, m.User.Username+m.User.Discriminator, time.Now().Local().String())
			cmd.Handler(s, m)
		}
	}
}
