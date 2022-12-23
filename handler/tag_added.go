package handler

import (
	"example-bot/api"
	"github.com/traPtitech/traq-ws-bot/payload"
	"log"
)

func TagAdded() func(p *payload.TagAdded) {
	return func(p *payload.TagAdded) {
		log.Println("=================================================")
		log.Println("TagAdded()")
		log.Printf("Payload:"+"%+v", p)

		api.PostMessage("9f551eae-0e50-4887-984e-ce9d8b3919cc", "タグが追加されました: "+p.Tag)
	}
}
