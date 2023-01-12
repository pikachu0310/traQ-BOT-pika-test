package api

import (
	"context"
	"example-bot/util"
	"fmt"
	"github.com/traPtitech/go-traq"
)

func AddStamps(messageID string, stampID string) {
	fmt.Println("AddStamps", messageID, stampID)
	bot := util.GetBot()
	_, err := bot.API().StampApi.AddMessageStamp(context.Background(), messageID, stampID).Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveStamp(messageID string, stampID string) {
	fmt.Println("RemoveStamps", messageID, stampID)
	bot := util.GetBot()
	_, err := bot.API().StampApi.RemoveMessageStamp(context.Background(), messageID, stampID).Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func GetStamps(messageID string) []traq.MessageStamp {
	fmt.Println("GetStamps", messageID)
	bot := util.GetBot()
	stamps, _, err := bot.API().StampApi.GetMessageStamps(context.Background(), messageID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return stamps
}

func GetAllStamps() []traq.Stamp {
	fmt.Println("GetAllStamps")
	bot := util.GetBot()
	stamps, _, err := bot.API().StampApi.GetStamps(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return stamps
}

func GetStamp(stampID string) *traq.Stamp {
	fmt.Println("GetStamp", stampID)
	bot := util.GetBot()
	stamp, _, err := bot.API().StampApi.GetStamp(context.Background(), stampID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return stamp
}

func AddStamp(stamp traq.Stamp) {
	fmt.Println("AddStamp", stamp)
	bot := util.GetBot()
	_, err := bot.API().StampApi.AddMessageStamp(context.Background(), stamp.Id, stamp.Id).Execute()
	if err != nil {
		fmt.Println(err)
	}
}
