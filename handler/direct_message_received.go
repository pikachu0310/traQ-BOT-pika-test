package handler

import (
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

		if slice[0] == "@BOT_pika_test" {
			slice = slice[1:]
		}

		CommandReceived(slice, p.Message.ID, p.Message.ChannelID)
	}
}
