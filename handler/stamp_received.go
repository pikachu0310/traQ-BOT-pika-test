package handler

import (
	"example-bot/api"
	"example-bot/commands"
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
		commands.OxGamePlay(p.MessageID, p.Stamps)

	}
}

func searchingStamp(stampID string, messageID string) {
	for _, searching := range commands.SearchingList {
		if searching.MessageID != messageID {
			continue
		}
		if commands.SearchingStamp == stampID {
			api.EditMessage(messageID, "正解です！")
		}
	}
}
