package hipchat

import (
	"fmt"
	"log"
	"time"

	"github.com/adams-sarah/go-xmpp"
)

//Client represents an XMPP client
type Client struct {
	url      string
	port     int
	username string
	password string
	Client   *xmpp.Client
	alive    bool
}

//NewClient creates a new client using xmpp.Options
//and set the bot status as available using a keep alive go routine
func NewClient(url string, port int, username, password string) *Client {
	c := &Client{
		url:      url,
		port:     port,
		username: username,
		password: password,
	}

	opts := xmpp.Options{
		Host:     fmt.Sprintf("%s:%d", c.url, c.port),
		User:     fmt.Sprintf("%s@%s", c.username, c.url),
		Password: c.password,
		Debug:    true,
		Resource: "bot",
	}
	// Initialize client
	hipchat, err := opts.NewClient()
	if err != nil {
		log.Println("Client error:", err)
	}

	c.Client = hipchat

	c.setAvailable()

	return c
}

func (c *Client) setAvailable() {
	c.alive = true
	go c.keepAlive()
}

func (c *Client) setOffline() {
	c.alive = false
}

//Send a presence message every 30 seconds
func (c *Client) keepAlive() {
	for c.alive {
		c.Client.SendOrg(" ")
		time.Sleep(30 * time.Second)
	}
}
