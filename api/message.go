package api

import (
	"context"
	"example-bot/util"
	"log"

	"github.com/traPtitech/go-traq"
)

func PostMessage(channelID string, content string) {

	bot := util.GetBot()

	_, _, err := bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		log.Println(err)
	}
}
