package commands

import "example-bot/api"

func BotJoin(args ArgsV2) {
	alreadyJoined, err := api.IsBotJoined(args.ChannelID)
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on IsBotJoined: "+err.Error())
		return
	}
	if alreadyJoined {
		api.PostMessage(args.ChannelID, "僕はここにいるよ？")
		return
	}
	err = api.BotJoin(args.ChannelID)
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on BotJoin: "+err.Error())
		return
	}
	api.PostMessage(args.ChannelID, "やっほー！僕は!{\"type\":\"user\",\"raw\":\"@BOT_pika_test\",\"id\":\"f60932e2-2a57-494f-87df-b7aa91b5c8b6\"}！よろしくだぞ！\n僕の能力や~~黒~~歴史について知りたかったら、[ここ](https://wiki.trap.jp/bot/pika_test)を見てね！ (/help でも見れるぞ！)")
}

func BotLeave(args ArgsV2) {
	alreadyJoined, err := api.IsBotJoined(args.ChannelID)
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on IsBotJoined: "+err.Error())
		return
	}
	if !alreadyJoined {
		api.PostMessage(args.ChannelID, "僕はここにいないよ？")
		return
	}
	err = api.BotLeave(args.ChannelID)
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on BotLeave: "+err.Error())
		return
	}
	api.PostMessage(args.ChannelID, "またね！")
}
