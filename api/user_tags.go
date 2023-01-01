package api

import (
	"context"
	"example-bot/util"
	"fmt"
	"github.com/traPtitech/go-traq"
)

func GetTags(userID string) []traq.UserTag {
	fmt.Println("GetUser", userID)
	bot := util.GetBot()
	UserTags, _, err := bot.API().UserTagApi.GetUserTags(context.Background(), userID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return UserTags
}

func GetTag(tagID string) *traq.Tag {
	fmt.Println("GetUser", tagID)
	bot := util.GetBot()
	UserTag, _, err := bot.API().UserTagApi.GetTag(context.Background(), tagID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return UserTag
}
