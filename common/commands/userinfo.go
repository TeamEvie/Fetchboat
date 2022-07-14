package commands

import "github.com/bwmarrin/discordgo"

func UserInfoHandler(s *discordgo.Session, m *discordgo.InteractionCreate) {
	s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "woosh",
		},
	})
}
