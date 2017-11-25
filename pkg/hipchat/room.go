package hipchat

import (
	"github.com/adams-sarah/go-xmpp"
)

//Room represent an XMPP Room (a group chat)
type Room struct {
	client   *xmpp.Client
	roomJid  string
	fullname string
	roomType string
}

//NewRoom creates a new room.
//parameters are :
//* The Jabber room ID
//* The name (fullname) used in the room.
func (c *Client) NewRoom(roomJid, fullname string) *Room {
	return &Room{
		client:   c.Client,
		roomJid:  roomJid,
		fullname: fullname,
		roomType: "groupchat",
	}
}

//Join makes the bot join the room using the fullname specified
func (r *Room) Join() {
	r.client.JoinMUC(r.roomJid, r.fullname)
}

//xmpp.Chat{To: roomJid, From: roomJid + "/" + fullname, Type: "groupchat", Text: message}
