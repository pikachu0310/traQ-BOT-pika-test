package commands

import (
	"fmt"
	"regexp"

	"example-bot/api"
)

func Edit(cmdText string) {
	args := CmdArgs(cmdText)
	if len(args) <= 1 {
		return
	}
	fmt.Println(cmdText)
	match := regexp.MustCompile(args[0] + " ")
	api.EditMessage(args[0], match.ReplaceAllString(cmdText, ""))
}
