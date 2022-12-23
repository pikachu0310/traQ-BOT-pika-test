package handler

import (
	"example-bot/api"
	"github.com/traPtitech/traq-ws-bot/payload"
	"log"
)

func TagRemoved() func(p *payload.TagRemoved) {
	return func(p *payload.TagRemoved) {
		log.Println("=================================================")
		log.Println("TagRemoved()")
		log.Printf("Payload:"+"%+v", p)

		api.PostMessage("9f551eae-0e50-4887-984e-ce9d8b3919cc", "タグが削除されました: "+p.Tag)
	}
}
