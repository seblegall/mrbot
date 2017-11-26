package hipchat

//Message represent an xmpp chat message
type Message struct {
	To   string
	From string
	Type string
	Text string
}

//NewMessage return a new Message
func NewMessage(to, from, messageType, text string) *Message {

	return &Message{
		To:   to,
		From: from,
		Type: messageType,
		Text: text,
	}
}
