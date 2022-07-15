package commands

import (
	"runtime"

	"github.com/bwmarrin/discordgo"
)

var (
	PingCommand = &Command{
		Handler: handler,
		Name:    "goping",
	}
)

func handler(s *discordgo.Session, m *discordgo.InteractionCreate) {
	s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "this came from go!\n\nGo Version: " + runtime.Version() + "\nGo OS: " + runtime.GOOS + "\nGo Arch: " + runtime.GOARCH + "\nDiscordgo Version: " + discordgo.VERSION,
		},
	})
}
