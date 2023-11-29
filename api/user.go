package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/traPtitech/go-traq"

	"example-bot/util"
)

var UserList []traq.User = GetUsers()
var UserAndBotList []traq.User = append(UserList, BotsToUsers(GetBots())...)

func GetUser(userID string) *traq.UserDetail {
	fmt.Println("GetUser", userID)
	bot := util.GetBot()
	User, _, err := bot.API().UserApi.GetUser(context.Background(), userID).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return User
}

func GetUsers() []traq.User {
	bot := util.GetBot()
	Users, _, err := bot.API().UserApi.GetUsers(context.Background()).IncludeSuspended(true).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return Users
}

func BotsToUsers(Bots []traq.Bot) []traq.User {
	var Users []traq.User
	for _, bot := range Bots {
		Users = append(Users, BotToUser(bot))
	}
	return Users
}

// GetUserByUserNameWithMe 一致するユーザーが見つからなかったら、自分を返す
func GetUserByUserNameWithMe(UserName string, UserID string) *traq.User {
	meNum := 0
	for num, user := range UserAndBotList {
		if user.Name == UserName {
			return &user
		}
		if user.Id == UserID {
			meNum += num
		}
	}
	return &UserAndBotList[meNum]
}

func GetUserByUserName(UserName string) (*traq.User, error) {
	for _, user := range UserAndBotList {
		if user.Name == UserName {
			return &user, nil
		}
	}
	err := errors.New("ユーザーが見つかりませんでした")
	return nil, err
}
