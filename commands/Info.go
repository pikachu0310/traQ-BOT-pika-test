package commands

import (
	"example-bot/api"
	"fmt"
)

func Info(args Args) {
	if len(args.Slice) <= 1 {
		return
	}
	if args.Slice[1] == "help" {
		api.PostMessage(args.ChannelID, "Usage: @BOT_pika_test /id :pikachu:")
		return
	}
	userInfo, err := UserInfo(args.Slice[1])
	if err == nil {
		api.PostMessage(args.ChannelID, userInfo)
		return
	}
	stampInfo, err := StampInfo(args.Slice[1])
	if err == nil {
		api.PostMessage(args.ChannelID, stampInfo)
		return
	}
	if len(args.Slice[1]) < 2 {
		notFound(args)
	}
	stampName := args.Slice[1][1 : len(args.Slice[1])-1]
	stampInfo, err = StampInfo(stampName)
	if err == nil {
		api.PostMessage(args.ChannelID, stampInfo)
		return
	}
	if len(args.Slice[1]) < 3 {
		notFound(args)
	}
	userName := args.Slice[1][2 : len(args.Slice[1])-1]
	userInfo, err = UserInfo(userName)
	if err == nil {
		api.PostMessage(args.ChannelID, userInfo)
		return
	}
	notFound(args)
}

func StampInfo(stampName string) (string, error) {
	stamp, err := api.GetStampByStampName(stampName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("スタンプUUID: %s\nスタンプ名: %s\n(pattern: ^[a-zA-Z0-9_-]{1,32}$)\n作成者UUID: %s\n作成日時: %s\n更新日時: %s\nファイルUUID: %s\nUnicode絵文字か: %s", stamp.Id, stamp.Id, stamp.CreatorId, stamp.CreatedAt, stamp.UpdatedAt, stamp.FileId, stamp.IsUnicode), nil
}

func UserInfo(userName string) (string, error) {
	user, err := api.GetUserByUserName(userName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("ユーザーUUID: %s\nユーザー名: %s\n(pattern: ^[a-zA-Z0-9_-]{1,32}$)\nユーザー表示名: %s\n(minLength: 0, maxLength: 32)\nアイコンファイルUUID: %s\nBOTかどうか: %s\nユーザーアカウント状態 0: 停止 1: 有効 2: 一時停止: %s\n更新日時: %s", user.Id, user.Name, user.DisplayName, user.IconFileId, user.Bot, user.State, user.UpdatedAt), nil
}

func notFound(args Args) {
	api.PostMessage(args.ChannelID, "一致するユーザーかスタンプが見つかりませんでした。")
}
