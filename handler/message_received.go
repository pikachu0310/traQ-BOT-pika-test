package handler

import (
	"example-bot/api"
	"example-bot/util"
	"fmt"
	"log"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func MessageReceived() func(p *payload.MessageCreated) {
	return func(p *payload.MessageCreated) {
		log.Println("=================================================")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:")
		log.Println(p.Message.Text)
		log.Println("Payload:")
		log.Printf("%+v\n", p)

		text := p.Message.PlainText
		slice := strings.Split(text, " ")

		if slice[0] == "/slice" {
			respond(p, strings.Join(slice, ", "))
		} else if slice[0] == "/ping" {
			respond(p, "pong")
		} else if slice[0] == "/oisu" {
			oisu_slice := []int{0, 1, 2, 3}
			oisu_str := []string{":oisu-1:", ":oisu-2:", ":oisu-3:", ":oisu-4yoko:"}
			util.Shuffle(oisu_slice)
			var oisu string = ""
			for i := 0; i < 4; i++ {
				oisu += fmt.Sprintf(oisu_str[oisu_slice[i]])
			}

			respond(p, oisu+" "+p.Message.User.DisplayName)
		}
	}
}

func respond(p *payload.MessageCreated, content string) {
	api.PostMessage(p.Message.ChannelID, content)
}
