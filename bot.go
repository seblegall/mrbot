package main

import (
	"fmt"

	"github.com/seblegall/mrbot/pkg/gitlab"
	"github.com/seblegall/mrbot/pkg/hipchat"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"log"
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
			fmt.Println("message received !")
			b.Answer(m)
		}
	}(b)
}


type answer struct {
	Result struct{
		Speech string `json:"speech"`
	} `json:"result"`
}

//Answer makes the bot respond to a given message.
//This is where answer rules are defined.
func (b *Bot) Answer(m *hipchat.Message) {

	fmt.Println("let's call the dialogflow API")
	url := "https://api.dialogflow.com/v1/query"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(fmt.Sprintf(`
	{
		"lang": "fr",
		"query": "%s",
		"sessionId": "12345"
	}
	`, m.Text))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer c00af3e17488431fa2a84272533f40ca")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Response of the api is : %s\n", string(body))

	var a answer

	e := json.Unmarshal([]byte(body), &a )

	if e != nil {
		log.Println(err)
	}

	b.room.Send(a.Result.Speech)
}
