package handler

import (
	"example-bot/api"
	"example-bot/commands"
	"log"
	"strconv"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

const prefix = "/"

type CommandFunc = func(args commands.Args)

var commandsMap map[string]CommandFunc

func init() {
	commandsMap = map[string]CommandFunc{
		"slice": func(args commands.Args) {
			respond(args.ChannelID, strings.Join(args.Slice, ", "))
		},
		"ping": func(args commands.Args) {
			respond(args.ChannelID, "pong")
		},
		"oisu": func(args commands.Args) {
			commands.Oisu(args.ChannelID)
		},
		"stamp": func(args commands.Args) {
			if len(args.Slice) <= 2 {
				return
			}
			verb := args.Slice[1]
			stampID := args.Slice[2]
			switch verb {
			case "add":
				api.AddStamps(args.MessageID, stampID)
			case "remove":
				api.RemoveStamp(args.MessageID, stampID)
			}
		},
		"allstamps": func(args commands.Args) {
			if len(args.Slice) == 1 {
				return
			}

			num := args.Slice[1]
			toInt, err := strconv.Atoi(num)
			if err != nil {
				return
			}

			allStamps := api.GetAllStamps()
			// const stampsRespond = allStamps.takeFirstN(num).map((stamp) => `:${stamp.Name}:`).join("")
			stampsRespond := ""
			for i := 0; i <= toInt; i++ {
				stampsRespond += ":" + allStamps[i].Name + ":"
			}
			respond(args.ChannelID, stampsRespond)
		},
		"game": func(args commands.Args) {
			commands.OxGameStart(args.ChannelID, args.Slice)
		},
		"edit": func(args commands.Args) {
			//api.EditMessage(p.Message.ID, slice[1])
			if len(args.Slice) == 3 {
				api.EditMessage(args.Slice[1], args.Slice[2])
			}
		},
		"help": func(args commands.Args) {
			commands.Help(args.ChannelID, args.Slice)
		},
		"sql": func(args commands.Args) {
			commands.Sql(args.ChannelID, args.Slice)
		},
		"tag": func(args commands.Args) {
			commands.Tag(args.ChannelID, args.UserID, args.Slice)
		},
		"docker": func(args commands.Args) {
			commands.Docker(args.ChannelID, args.Slice)
		},
		//"stamps": func(args commands.Args) {
		//	commands.Stamps(args)
		//},
	}
}

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

	// check prefix
	commandName := slice[0]
	if !strings.HasPrefix(commandName, prefix) {
		return
	}
	commandName = commandName[len(prefix):] // strip prefix

	if cmd, ok := commandsMap[commandName]; ok {
		cmd(commands.Args{
			Slice:     slice,
			MessageID: MessageID,
			ChannelID: ChannelID,
			UserID:    UserID,
		})
	}
}

func respond(ChannelID, content string) {
	api.PostMessage(ChannelID, content)
}
