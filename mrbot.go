package main

import (
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
	password    = "xcem7sa2"
	roomJid     = "175921_testbot@conf.hipchat.com"
)

func main() {

	hipchat := hipchat.NewClient(hipChatJabberURL, hipChatJabberPort, username, password)
	room := hipchat.NewRoom(roomJid, fullname)
	bot := NewBot(hipchat, room)
	bot.Join()
	bot.ListenAndAnswer()
}
