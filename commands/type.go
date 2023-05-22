package commands

import "strings"

type Args struct {
	Slice     []string
	MessageID string
	ChannelID string
	UserID    string
}

type ArgsV2 struct {
	MessageText  string
	MessageID    string
	ChannelID    string
	UserID       string
	OriginalText string
}

func CmdArgs(text string) []string {
	slice := strings.Split(text, " ")
	if slice[0] == "@BOT_pika_test" {
		slice = slice[1:]
	}
	return slice
}
