package main

import (
	"example-bot/handler"
	"example-bot/util"
)

func main() {

	handler.Cron()

	bot := util.GetBot()

	bot.OnMessageCreated(handler.MessageReceived())
	bot.OnDirectMessageCreated(handler.DirectMessageReceived())
	bot.OnChannelTopicChanged(handler.ChannelTopicChanged())
	bot.OnTagAdded(handler.TagAdded())
	bot.OnTagRemoved(handler.TagRemoved())
	bot.OnBotMessageStampsUpdated(handler.StampReceived())
	// bot.OnDirectMessageCreated(handler.MessageReceived())

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
