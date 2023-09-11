package bot

import (
	"dc-bot/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
	discord *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}

	u, err := goBot.User("@me")
	if err != nil {
		panic(err)
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is runnin!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}
}