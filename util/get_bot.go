package util

import (
	"fmt"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func GetBot() (bot *traqwsbot.Bot) {
	token := GetToken()

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: token,
	})
	if err != nil {
		fmt.Printf("error: Bot変数が作れなかった!: %v", err)
	}
	return bot
}
