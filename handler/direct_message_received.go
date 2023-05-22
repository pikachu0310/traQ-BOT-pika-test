package handler

import (
	"log"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func DirectMessageReceived() func(p *payload.DirectMessageCreated) {
	return func(p *payload.DirectMessageCreated) {
		log.Println("=================================================")
		log.Printf("DirectMessageReceived()")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:" + p.Message.Text)
		log.Printf("Payload:"+"%+v", p)

		if p.Message.User.Bot {
			return
		}

		CommandReceived(p.Message.PlainText, p.Message.ID, p.Message.ChannelID, p.Message.User.ID, p.Message.Text)
	}
}
