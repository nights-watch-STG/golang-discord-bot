package main

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func NewBot() (*discordgo.Session, error) {
	goBot, err := discordgo.New("Bot " + os.Getenv("MyToken"))
	return goBot, err
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	if strings.ToUpper(m.Content) == "LTV" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "tu LTV panchooda")
	}
	if m.ContentWithMentionsReplaced() == `<@&998929334645563502>` { //@fudi deya
		_, _ = s.ChannelMessageSend(m.ChannelID, "kithe marre aa kuriyave sare")
	}
	if m.ContentWithMentionsReplaced() == `<@&765554608172957696>` { //@kings
		_, _ = s.ChannelMessageSend(m.ChannelID, "Acha, gusse sare king bnate aa")
	}
	if m.ContentWithMentionsReplaced() == `<@&789483612072181771>` { //@ace
		_, _ = s.ChannelMessageSend(m.ChannelID, "aces te oda he FUDU aa")
	}
	if m.ContentWithMentionsReplaced() == `<@&1090312377876107324>` { //@shady sandhu sunyara squad
		_, _ = s.ChannelMessageSend(m.ChannelID, "serious kaamm aa jalde kathe hovo")
	}

	if m.ContentWithMentionsReplaced() == `@Sahil Sandhu` { // personal
		_, _ = s.ChannelMessageSend(m.ChannelID, "peo nu kato vaja marre janda")
	}
}
