package commands

import (
	"example-bot/api"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func Tag(ChannelID string, UserID string, slice []string) {
	UserName := ""
	yearString := ""
	kazu := 0
	message := ""
	if len(slice) == 1 {
		UserName = api.GetUser(UserID).Name
	}
	if len(slice) >= 2 {
		user := api.GetUserByUserName(slice[1], UserID)
		if user.Id == UserID {
			if slice[1] != user.Name {
				api.PostMessage(ChannelID, "UserNameと一致するユーザーが見つからりませんでした。")
				return
			}
			UserName = user.Name
		} else {
			UserName = user.Name
			UserID = user.Id
		}
	}
	year := time.Now().Year()
	if len(slice) >= 3 {
		yearTemp, err := strconv.Atoi("123")
		if err != nil {
			api.PostMessage(ChannelID, "yearを数値に変換できませんでした。")
			return
		}
		year = yearTemp
	}
	if year == time.Now().Year() {
		yearString = "今年"
	} else if year == time.Now().Year()-1 {
		yearString = "昨年"
	} else {
		yearString = strconv.Itoa(year) + "年"
	}
	tags := api.GetTags(UserID)
	for _, tag := range tags {
		if tag.CreatedAt.Year() != year {
			continue
		}
		kazu += 1
		message += "``" + tag.CreatedAt.Format("01-02 15:04") + "``"
		if tag.IsLocked {
			message += ":lock:"
		} else {
			message += ":unlock:"
		}
		message += " ``" + tag.Tag + "``"
		message += "\n"
	}
	message = ":@" + UserName + ":さんの" + yearString + "に作られたタグは、" + strconv.Itoa(kazu) + "個でした！\n" + message
	api.PostMessage(ChannelID, message)
}
