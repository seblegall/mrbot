package hipchat

import (
	"log"
	"strings"

	xmpp "github.com/adams-sarah/go-xmpp"
)

//Stream reprensents a chat message stream.
type Stream struct {
	client *xmpp.Client
	C      chan xmpp.Chat
	run    bool
}

func (c *Client) newStream(mentionname string) *Stream {
	stream := Stream{
		client: c.Client,
		C:      make(chan xmpp.Chat),
	}

	stream.start(mentionname)
	return &stream
}

func (s *Stream) start(mentionname string) {
	s.run = true
	go s.loop(mentionname)
}

func (s *Stream) loop(mentionname string) {
	defer close(s.C)
	for s.run {
		message, err := s.client.Recv()
		if err != nil {
			log.Fatal(err)
		}

		if chatMsg, ok := message.(xmpp.Chat); ok {
			if strings.HasPrefix(chatMsg.Text, "@"+mentionname) {
				s.C <- chatMsg
			}
		}
	}
}

//Stream start a new stream that listen message containing
//a reference (using "@") to the mentionname specified
func (c *Client) Stream(mentionname string) (stream *Stream) {
	return c.newStream(mentionname)
}
