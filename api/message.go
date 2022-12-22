package api

import (
	"context"
	"example-bot/util"
	"fmt"
	"log"

	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func PostMessage(channelID string, content string) {

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

	_, _, err = bot.API().
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
