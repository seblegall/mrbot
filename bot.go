package main

import (
	"fmt"
	"strings"

	"github.com/seblegall/mrbot/pkg/gitlab"
	"github.com/seblegall/mrbot/pkg/hipchat"
)

//Bot is a robot
type Bot struct {
	hipchat *hipchat.Client
	room    *hipchat.Room
	gitlab  *gitlab.Client
}

//NewBot creates a new bot using an hipchat client and set a room for the bot to join.
func NewBot(client *hipchat.Client, room *hipchat.Room, gitlab *gitlab.Client) *Bot {
	bot := &Bot{
		hipchat: client,
		room:    room,
		gitlab:  gitlab,
	}

	return bot
}

//Join make the bot join his room.
func (b *Bot) Join() {
	b.room.Join()
}

//ListenAndAnswer make the bot listen for message that mention (using "@")
//the bot and try to respond to it
func (b *Bot) ListenAndAnswer() {

	go func(b *Bot) {
		stream := b.hipchat.Stream(mentionname)

		for m := range stream.C {
			b.Answer(m)
		}
	}(b)
}

//Answer makes the bot respond to a given message.
//This is where answer rules are defined.
func (b *Bot) Answer(m *hipchat.Message) {

	switch {
	case strings.Contains(m.Text, "coucou"):
		b.room.Send(fmt.Sprintf("Coucou %s !", m.From))

	default:
		b.room.Send("Je ne suis pas programmé pour répondre à cela.")
	}
}
