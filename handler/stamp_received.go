package handler

import (
	"fmt"
	"github.com/traPtitech/traq-ws-bot/payload"
	"log"
)

func StampReceived() func(p *payload.BotMessageStampsUpdated) {
	return func(p *payload.BotMessageStampsUpdated) {
		log.Println("=================================================")
		log.Println("StampReceived()")
		log.Printf("Payload:"+"%+v", p)

		fmt.Println(p.Stamps, p.MessageID)
	}
}
