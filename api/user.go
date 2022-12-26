package api

import (
	"context"
	"example-bot/util"
	"fmt"
	"github.com/traPtitech/go-traq"
)

func GetUser(userID string) *traq.UserDetail {
	fmt.Println("GetUser", userID)
	bot := util.GetBot()
	User, _, err := bot.API().UserApi.GetUser(context.Background(), userID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return User
}
