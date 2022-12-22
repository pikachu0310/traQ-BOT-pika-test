package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	godotenv "github.com/joho/godotenv"
	traq "github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	payload "github.com/traPtitech/traq-ws-bot/payload"
)

func GetToken() (token string, err error) {
	err = godotenv.Load(".env")
	token = os.Getenv("token")
	return token, err
}

func Shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {

	token, err := GetToken()

	if err != nil {
		fmt.Printf("error: tokenが読み込めなかった!: %v", err)
	}

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: token,
	})
	if err != nil {
		panic(err)
	}

	bot.OnMessageCreated(func(p *payload.MessageCreated) {

		log.Println("=================================================")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:")
		log.Println(p.Message.Text)
		log.Println("Payload:")
		log.Printf("%+v\n", p)

		text := p.Message.PlainText
		slice := strings.Split(text, " ")

		if slice[0] == "/slice" {

			log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
			_, _, err := bot.API().
				MessageApi.
				PostMessage(context.Background(), p.Message.ChannelID).
				PostMessageRequest(traq.PostMessageRequest{
					Content: string(strings.Join(slice, ", ")),
				}).
				Execute()
			if err != nil {
				log.Println(err)
			}
		} else if slice[0] == "/ping" {
			log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
			_, _, err := bot.API().
				MessageApi.
				PostMessage(context.Background(), p.Message.ChannelID).
				PostMessageRequest(traq.PostMessageRequest{
					Content: "pong",
				}).
				Execute()
			if err != nil {
				log.Println(err)
			}
		} else if slice[0] == "/oisu" {
			oisu_slice := []int{0, 1, 2, 3}
			oisu_str := []string{":oisu-1:", ":oisu-2:", ":oisu-3:", ":oisu-4yoko:"}
			Shuffle(oisu_slice)
			var oisu string = ""
			for i := 0; i < 4; i++ {
				oisu += fmt.Sprintf(oisu_str[oisu_slice[i]])
			}

			log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
			_, _, err := bot.API().
				MessageApi.
				PostMessage(context.Background(), p.Message.ChannelID).
				PostMessageRequest(traq.PostMessageRequest{
					Content: oisu + " " + p.Message.User.DisplayName,
				}).
				Execute()
			if err != nil {
				log.Println(err)
			}
		}
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
