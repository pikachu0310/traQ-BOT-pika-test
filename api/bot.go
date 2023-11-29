package api

import (
	"context"
	"example-bot/util"
	"fmt"
	"github.com/traPtitech/go-traq"
)

const BotID = "28de7741-6c36-4c53-94e0-4ace05c221b5"

func BotJoin(ChannelID string) error {
	bot := util.GetBot()
	_, err := bot.API().BotApi.LetBotJoinChannel(context.Background(), BotID).PostBotActionJoinRequest(traq.PostBotActionJoinRequest{ChannelId: ChannelID}).Execute()
	return err
}

func BotLeave(ChannelID string) error {
	bot := util.GetBot()
	_, err := bot.API().BotApi.LetBotLeaveChannel(context.Background(), BotID).PostBotActionLeaveRequest(traq.PostBotActionLeaveRequest{ChannelId: ChannelID}).Execute()
	return err
}

func IsBotJoined(ChannelID string) (bool, error) {
	bot := util.GetBot()
	bots, _, err := bot.API().BotApi.GetChannelBots(context.Background(), ChannelID).Execute()
	if err != nil {
		return false, err
	}
	for _, bot := range bots {
		if bot.Id == BotID {
			return true, nil
		}
	}
	return false, nil
}

func GetBots() []traq.Bot {
	bot := util.GetBot()
	Bots, _, err := bot.API().BotApi.GetBots(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return Bots
}

func BotToUser(bot traq.Bot) traq.User {
	user := traq.User{
		Id:          bot.Id,
		Name:        bot.BotUserId,
		DisplayName: "",
		IconFileId:  "",
		Bot:         true,
		State:       BotStateToUserState(bot.State),
		UpdatedAt:   bot.UpdatedAt,
	}
	return user
}

func BotStateToUserState(botState traq.BotState) traq.UserAccountState {
	switch botState {
	case traq.BOTSTATE_deactivated:
		return traq.USERACCOUNTSTATE_deactivated
	case traq.BOTSTATE_active:
		return traq.USERACCOUNTSTATE_active
	}
	return traq.USERACCOUNTSTATE_suspended
}
