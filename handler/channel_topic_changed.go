package handler

import (
	"example-bot/api"
	"github.com/traPtitech/traq-ws-bot/payload"
	"log"
)

func ChannelTopicChanged() func(p *payload.ChannelTopicChanged) {
	return func(p *payload.ChannelTopicChanged) {
		log.Println("=================================================")
		log.Printf("ChannelTopicChanged()")
		log.Printf("Payload:"+"%+v", p)

		content := "channel topicが変更されました: " + p.Topic
		api.PostMessage(p.Channel.ID, content)
	}
}
