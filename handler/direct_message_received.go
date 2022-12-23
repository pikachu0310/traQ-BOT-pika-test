package handler

import (
	"example-bot/api"
	"example-bot/util"
	"fmt"
	"log"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func DirectMessageReceived() func(p *payload.DirectMessageCreated) {
	return func(p *payload.DirectMessageCreated) {
		log.Println("=================================================")
		log.Printf("DirectMessageReceived()")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:" + p.Message.Text)
		log.Printf("Payload:"+"%+v", p)

		text := p.Message.PlainText
		slice := strings.Split(text, " ")

		if slice[0] == "/slice" {
			direct_respond(p, strings.Join(slice, ", "))
		} else if slice[0] == "/ping" {
			direct_respond(p, "pong")
		} else if slice[0] == "/oisu" {
			oisu_slice := []int{0, 1, 2, 3}
			oisu_str := []string{":oisu-1:", ":oisu-2:", ":oisu-3:", ":oisu-4yoko:"}
			util.Shuffle(oisu_slice)
			var oisu string = ""
			for i := 0; i < 4; i++ {
				oisu += fmt.Sprintf(oisu_str[oisu_slice[i]])
			}
			direct_respond(p, oisu+" "+p.Message.User.DisplayName)
		} else if slice[0] == "/help" {
			direct_respond(p, "そんなコマンドはないよ")
		} else if slice[0] == "/stamp" {
			api.AddStamps(p.Message.ID, slice[1])
		}
	}
}

func direct_respond(p *payload.DirectMessageCreated, content string) {
	api.PostMessage(p.Message.ChannelID, content)
}
