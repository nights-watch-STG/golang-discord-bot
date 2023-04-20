package main

import "github.com/bwmarrin/discordgo"

func HandleBot() {
	Bot.AddHandler(messageHandler)
	Bot.Identify.Intents = discordgo.IntentsGuildMessages
	err := Bot.Open()
	if err != nil {
		return
	}
}
