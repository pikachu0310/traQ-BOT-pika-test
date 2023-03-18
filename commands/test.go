package commands

import (
	"example-bot/api"
)

func GetTitle(args ArgsV2) {
	slice := StringToSlice(args.MessageText)
	if len(slice) < 2 {
		return
	}
	api.PostMessage(args.ChannelID, api.GetTitle(slice[1]))
}
