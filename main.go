package main

import (
	"example-bot/handler"
	"example-bot/util"
	"fmt"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func main() {

	token, err := util.GetToken()

	if err != nil {
		fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	}

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: token,
	})
	if err != nil {
		panic(err)
	}

	bot.OnMessageCreated(handler.MessageReceived())
	bot.OnDirectMessageCreated(handler.DirectMessageReceived())
	bot.OnChannelTopicChanged(handler.ChannelTopicChanged())
	// bot.OnDirectMessageCreated(handler.MessageReceived())

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
