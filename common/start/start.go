package start

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TeamEvie/Fetchboat/common/commands"
	"github.com/bwmarrin/discordgo"
)

func Start() {

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		fmt.Println("Failed to create a session!", err)
		return
	}

	discord.AddHandler(messageCreate)
	discord.AddHandler(commands.UserInfoHandler)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Println("Failed to connect to the gateway!", err)
		return
	}

	fmt.Println("Started! - Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == s.State.User.Mention() {
		s.ChannelMessageSend(m.ChannelID, "alive!")
	}
}
