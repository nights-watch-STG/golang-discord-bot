package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func NewBot() (*discordgo.Session, error) {
	goBot, err := discordgo.New("Bot " + os.Getenv("MyToken"))
	return goBot, err
}

func HandleBot() {
	Bot.AddHandler(messageHandler)
	Bot.Identify.Intents = discordgo.IntentsGuildMessages
	err := Bot.Open()
	if err != nil {
		return
	}
}
