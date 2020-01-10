package discord

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func GetDiscordGo() *discordgo.Session {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	return s
}

func SendMessage(s *discordgo.Session, message string) {
	s.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL"), message)
}
