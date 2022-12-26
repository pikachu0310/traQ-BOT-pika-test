package api

import (
	"context"
	"example-bot/util"
	"log"

	"github.com/traPtitech/go-traq"
)

func PostMessage(channelID string, content string) *traq.Message {

	bot := util.GetBot()

	message, _, err := bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return message
}

func EditMessage(messageID string, content string) {

	bot := util.GetBot()

	_, err := bot.API().
		MessageApi.EditMessage(context.Background(), messageID).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	if err != nil {
		log.Println(err)
	}
}

func GetMessage(messageID string) *traq.Message {

	bot := util.GetBot()

	message, _, err := bot.API().
		MessageApi.
		GetMessage(context.Background(), messageID).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return message
}
