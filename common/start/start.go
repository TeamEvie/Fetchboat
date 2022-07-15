package start

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TeamEvie/Fetchboat/commands"
	"github.com/bwmarrin/discordgo"
)

func Start() {

	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		fmt.Println("Failed to create a session!", err)
		return
	}

	commands.Inject(s)

	err = s.Open()
	if err != nil {
		fmt.Println("Failed to connect to the gateway!", err)
		return
	}

	fmt.Println("Started! - Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	s.Close()
}
