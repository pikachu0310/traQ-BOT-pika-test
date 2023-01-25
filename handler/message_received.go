package handler

import (
	"example-bot/api"
	"example-bot/commands"
	"log"
	"strconv"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func MessageReceived() func(p *payload.MessageCreated) {
	return func(p *payload.MessageCreated) {
		log.Println("=================================================")
		log.Printf("MessageReceived()")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:" + p.Message.Text)
		log.Printf("Payload:"+"%+v", p)

		text := p.Message.PlainText
		slice := strings.Split(text, " ")

		if slice[0] == "@BOT_pika_test" {
			slice = slice[1:]
		}

		CommandReceived(slice, p.Message.ID, p.Message.ChannelID, p.Message.User.ID)
	}
}

func CommandReceived(slice []string, MessageID string, ChannelID string, UserID string) {
	if len(slice) == 0 {
		return
	}
	if slice[0] == "/slice" {
		respond(ChannelID, strings.Join(slice, ", "))
	} else if slice[0] == "/ping" {
		respond(ChannelID, "pong")
	} else if slice[0] == "/oisu" {
		commands.Oisu(ChannelID)
	} else if slice[0] == "/stamp" {
		if len(slice) == 1 {
			return
		}
		if slice[1] == "add" {
			api.AddStamps(MessageID, slice[1])
		} else if slice[1] == "remove" {
			api.RemoveStamp(MessageID, slice[1])
		}
	} else if slice[0] == "/allstamps" {
		if len(slice) == 1 {
			return
		}
		allStamps := api.GetAllStamps()
		stampsRespond := ""
		num := slice[1]
		toInt, err := strconv.Atoi(num)
		if err != nil {
		} else {
			for i := 0; i <= toInt; i++ {
				stampsRespond += ":" + allStamps[i].Name + ":"
			}
			respond(ChannelID, stampsRespond)
		}
	} else if slice[0] == "/game" {
		commands.OxGameStart(ChannelID, slice)
	} else if slice[0] == "/edit" {
		//api.EditMessage(p.Message.ID, slice[1])
		if len(slice) == 3 {
			api.EditMessage(slice[1], slice[2])
		}
	} else if slice[0] == "/help" {
		commands.Help(ChannelID, slice)
	} else if slice[0] == "/sql" {
		commands.Sql(ChannelID, slice)
	} else if slice[0] == "/tag" {
		commands.Tag(ChannelID, UserID, slice)
	} else if slice[0] == "/docker" {
		commands.Docker(ChannelID, slice)
	}

}

func respond(ChannelID, content string) {
	api.PostMessage(ChannelID, content)
}
