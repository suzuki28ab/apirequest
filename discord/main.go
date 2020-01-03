package discord

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func GetDiscordGo() *discordgo.Session {
	s, err := discordgo.New("Bot " + os.Getenv("YASUTAKA_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	return s
}

func SendMessage(s *discordgo.Session, str string) {
	s.ChannelMessageSend(os.Getenv("DAREGIRU_GENERAL"), str)
}
