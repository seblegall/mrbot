package main

import (
	"fmt"
	"os"
	"strings"

	xmpp "github.com/adams-sarah/go-xmpp"
	"github.com/seblegall/mrbot/pkg/hipchat"
)

const (
	// HipChat jabber info
	hipChatJabberURL  = "chat.hipchat.com"
	hipChatJabberPort = 5223
)

var (
	resource = "bot" // Kind of Hipchat user (probably shouldn't change this)

	// Vars needed for Hipbot to ping Hipchat:
	username    = "175921_5350262"
	mentionname = "mrbot"
	fullname    = "Mr Bot"
	password    = os.Getenv("PASSWORD")
	roomJid     = "175921_testbot@conf.hipchat.com"
)

var bot *hipchat.Client
var room *hipchat.Room

func main() {

	bot = hipchat.NewClient(hipChatJabberURL, hipChatJabberPort, username, password)

	room := bot.NewRoom(roomJid, fullname)
	room.Join()

	listen(bot.Stream(mentionname))
}

func listen(stream *hipchat.Stream) {

	for chatMsg := range stream.C {
		switch {
		case strings.Contains(chatMsg.Text, "coucou"):
			send(fmt.Sprintf("Coucou %s !", chatMsg.From))

		default:
			send("Je ne suis pas programmé pour répondre à cela.")
		}
	}

}

func send(message string) {

	bot.Client.Send(xmpp.Chat{To: roomJid, From: fullname, Type: "groupchat", Text: message})
}
