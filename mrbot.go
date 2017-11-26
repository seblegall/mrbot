package main

import (
	"fmt"

	"github.com/seblegall/mrbot/pkg/hipchat"
	"github.com/spf13/viper"
)

var (
	//Hipchat
	hipChatJabberURL  string
	hipChatJabberPort int
	username          string
	mentionname       string
	fullname          string
	password          string
	roomJid           string

	//Gitlab
	token     string
	gitlabURL string
	groups    []string
)

func main() {
	setConfig()

	hipchat := hipchat.NewClient(hipChatJabberURL, hipChatJabberPort, username, password)
	room := hipchat.NewRoom(roomJid, fullname)
	bot := NewBot(hipchat, room)
	bot.Join()
	bot.ListenAndAnswer()
	bot.ListenMergeRequest()
}

func setConfig() {
	viper.SetConfigName(".mrbot")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	//Hipchat configuration
	hipChatJabberURL = viper.GetString("hipchat.server")
	hipChatJabberPort = viper.GetInt("hipchat.port")
	username = viper.GetString("hipchat.username")
	mentionname = viper.GetString("hipchat.mentionname")
	fullname = viper.GetString("hipchat.fullname")
	password = viper.GetString("hipchat.password")
	roomJid = viper.GetString("hipchat.roomJid")

	//Gitlab configuration
	token = viper.GetString("gitlab.token")
	gitlabURL = viper.GetString("gitlab.url")
	groups = viper.GetStringSlice("gitlab.groups")

}
