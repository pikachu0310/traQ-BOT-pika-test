package handler

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"example-bot/api"
	"example-bot/commands"

	"github.com/traPtitech/traq-ws-bot/payload"
)

const prefix = "/"

type CommandFunc = func(args commands.Args)
type CommandFuncV2 = func(args commands.ArgsV2)

var commandsMap map[string]CommandFunc
var commandsMapV2 map[string]CommandFuncV2

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
			// api.EditMessage(p.Message.ID, slice[1])
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
		// "stamps": func(args commands.Args) {
		//	commands.Stamps(args)
		// },
		// "search": func(args commands.Args) {
		//	commands.Search(args)
		// },
		// "s": func(args commands.Args) {
		//	commands.Search(args)
		// },
		"info": func(args commands.Args) {
			commands.Info(args)
		},
		"id": func(args commands.Args) {
			commands.Info(args)
		},
	}

	commandsMapV2 = map[string]CommandFuncV2{
		"test": func(args commands.ArgsV2) {
			commands.GetTitle(args)
		},
	}
}

func MessageReceived() func(p *payload.MessageCreated) {
	return func(p *payload.MessageCreated) {
		log.Println("=================================================")
		log.Printf("MessageReceived()")
		log.Printf("Message created by %s\n", p.Message.User.DisplayName)
		log.Println("Message:" + p.Message.Text)
		log.Printf("Payload:"+"%+v", p)

		if p.Message.User.Bot {
			return
		}

		CommandReceived(p.Message.PlainText, p.Message.ID, p.Message.ChannelID, p.Message.User.ID)
	}
}

func CommandReceived(text, MessageID string, ChannelID string, UserID string) {

	args := commands.ArgsV2{MessageText: text, MessageID: MessageID, ChannelID: ChannelID, UserID: UserID}
	commandsV2(args)

	slice := strings.Split(text, " ")

	if slice[0] == "@BOT_pika_test" {
		slice = slice[1:]
	}

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

	if cmd, ok := commandsMapV2[commandName]; ok {
		cmd(args)
	}
}

func commandsV2(args commands.ArgsV2) {

	joinStampMatch := regexp.MustCompile(`join|いらっしゃい|oisu-|:oisu-1::oisu-2::oisu-3::oisu-4yoko:|おいすー`)
	leaveStampMatch := regexp.MustCompile(`leave|さようなら|:wave:|ばいばい`)
	mentionMatch := regexp.MustCompile(`@BOT_pika_test |@BOT_pika_test`)
	if mentionMatch.MatchString(args.MessageText) {
		if joinStampMatch.MatchString(args.MessageText) {
			commands.BotJoin(args)
		} else if leaveStampMatch.MatchString(args.MessageText) {
			commands.BotLeave(args)
		}
	}
	args.MessageText = mentionMatch.ReplaceAllString(args.MessageText, "")

	magStampMatch := regexp.MustCompile(`:(mag(|_right)|Internet_Explorer|google_g|yahoo)(\.[a-zA-Z_-]+)*:`)
	if magStampMatch.MatchString(args.MessageText) {
		textForSearch := magStampMatch.ReplaceAllString(args.MessageText, "")
		args.MessageText = textForSearch
		commands.Search(args)
		return
	}

	chatgptResetMatch := regexp.MustCompile(`(new(chat|gpt|chatgpt))|((chat|gpt|chatgpt)new)|(new (chat|gpt|chatgpt))|((chat|gpt|chatgpt) new)`)
	chatgptDebugMatch := regexp.MustCompile(`\/gptdebug|\/gpt debug`)
	chatgptChangeFirstSystemMessageMatch := regexp.MustCompile(`\/gptsystem|\/gptsys|\/gpt system|\/gpt sys`)
	chatgptShowFirstSystemMessageMatch := regexp.MustCompile(`show`)
	chatgptMatch := regexp.MustCompile(`(\/chat)|(\/chatgpt)|(\/gpt)|(:chatgpt(\.[a-zA-Z_-]+)*:)`)
	chatgpt4Match := regexp.MustCompile(`\/gpt4`)
	if chatgptMatch.MatchString(args.MessageText) {
		if chatgptResetMatch.MatchString(args.MessageText) {
			commands.ChatGPTReset(args)
			return
		} else if chatgptDebugMatch.MatchString(args.MessageText) {
			commands.ChatGPTDebug(args)
			return
		} else if chatgptChangeFirstSystemMessageMatch.MatchString(args.MessageText) {
			if chatgptShowFirstSystemMessageMatch.MatchString(args.MessageText) {
				commands.ChatGPTShowFirstSystemMessage(args)
				return
			}
			args.MessageText = chatgptChangeFirstSystemMessageMatch.ReplaceAllString(args.MessageText, "")
			commands.ChatGPTChangeFirstSystemMessage(args)
			return
		} else if chatgpt4Match.MatchString(args.MessageText) {
			textForSearch := chatgpt4Match.ReplaceAllString(args.MessageText, "")
			args.MessageText = textForSearch
			commands.ChatGPT4(args)
			return
		}
		textForSearch := chatgptMatch.ReplaceAllString(args.MessageText, "")
		args.MessageText = textForSearch
		commands.ChatGPT(args)
		return
	}

	cmdStampsDayMatch := regexp.MustCompile(`((\/kusa|\/stamps)+( |)+(day|today|period|daily)+( |))|((\/today|\/day|\/daily)+( |)+(|stamps|kusa)+( |))`)
	if cmdStampsDayMatch.MatchString(args.MessageText) {
		commands.StampsDay(cmdStampsDayMatch.ReplaceAllString(args.MessageText, ""), args.ChannelID)
		return
	}

	cmdStampsMatch := regexp.MustCompile(`\/kusa |\/stamps |\/kusa|\/stamps`)
	if cmdStampsMatch.MatchString(args.MessageText) {
		commands.Stamps(cmdStampsMatch.ReplaceAllString(args.MessageText, ""), args.ChannelID)
		return
	}

	cmdDeleteMatch := regexp.MustCompile(`\/delete `)
	if cmdDeleteMatch.MatchString(args.MessageText) {
		commands.Delete(cmdDeleteMatch.ReplaceAllString(args.MessageText, ""))
		return
	}
}

func respond(ChannelID, content string) {
	api.PostMessage(ChannelID, content)
}
