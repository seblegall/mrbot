package main

import (
	"github.com/seblegall/mrbot/pkg/hipchat"
	"github.com/seblegall/mrbot/pkg/dialogflow"
)

//Bot is a robot
type Bot struct {
	hipchat *hipchat.Client
	room    *hipchat.Room
	dialog *dialogflow.Client
}

//NewBot creates a new bot using an hipchat client and set a room for the bot to join.
func NewBot(client *hipchat.Client, room *hipchat.Room, dialog *dialogflow.Client) *Bot {
	bot := &Bot{
		hipchat: client,
		room:    room,
		dialog: dialog,
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
	b.room.Send(b.dialog.Query(string(m.Text)))
}
