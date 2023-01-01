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

func GetUsers() []traq.User {
	bot := util.GetBot()
	Users, _, err := bot.API().UserApi.GetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Println(err)
	}
	return Users
}

// GetUserByUserName 一致するユーザーが見つからなかったら、自分を返す
func GetUserByUserName(UserName string, UserID string) *traq.User {
	UserList := GetUsers()
	meNum := 0
	for num, user := range UserList {
		if user.Name == UserName {
			return &user
		}
		if user.Id == UserID {
			meNum += num
		}
	}
	return &UserList[meNum]
}
