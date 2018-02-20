package main

import (
	"os"
	"os/signal"
	"sync"

	"github.com/seblegall/mrbot/pkg/gitlab"
	"github.com/seblegall/mrbot/pkg/hipchat"
	"github.com/spf13/viper"
	"github.com/seblegall/mrbot/pkg/dialogflow"
	"strings"
)


const (
	hipChatJabberURL  = "chat.hipchat.com"
	hipChatJabberPort = 5223
	fullname = "Mr Bot"
	mentionname = "mrbot"
)

var (
	//Hipchat
	username          string
	password          string
	roomJid           string

	//Gitlab
	gitlabToken string
	gitlabURL   string
	groups      []string

	//Dialogflow
	dialogToken string
)

func main() {
	setConfig()

	hipchat := hipchat.NewClient(hipChatJabberURL, hipChatJabberPort, username, password)
	room := hipchat.NewRoom(roomJid, fullname)
	gitlab := gitlab.NewClient(gitlabURL, gitlabToken)
	dialog := dialogflow.NewClient(dialogToken)
	bot := NewBot(hipchat, room, gitlab, dialog)
	bot.Join()
	bot.ListenAndAnswer()
	//bot.ListenMergeRequest(groups)
	waitForCtrlC()
}

func setConfig() {
	//hipchat configuration
	viper.SetEnvPrefix("mrbot")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//Hipchat configuration
	username = viper.GetString("hipchat.username")
	password = viper.GetString("hipchat.password")
	roomJid = viper.GetString("hipchat.roomJid")

	//Gitlab configuration
	gitlabToken = viper.GetString("gitlab.token")
	gitlabURL = viper.GetString("gitlab.url")
	groups = viper.GetStringSlice("gitlab.groups")

	//Dialogflow configuration
	dialogToken = viper.GetString("dialogflow.token")

}

func waitForCtrlC() {
	var wg sync.WaitGroup
	wg.Add(1)
	var sig chan os.Signal

	sig = make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		wg.Done()
	}()
	wg.Wait()
}
