package main

import (
	"context"
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	traq "github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	payload "github.com/traPtitech/traq-ws-bot/payload"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	token := os.Getenv("token")

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: token,
	})
	if err != nil {
		panic(err)
	}
	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
		_, _, err := bot.API().
			MessageApi.
			PostMessage(context.Background(), p.Message.ChannelID).
			PostMessageRequest(traq.PostMessageRequest{
				Content: "oisu-",
			}).
			Execute()
		if err != nil {
			log.Println(err)
		}
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
