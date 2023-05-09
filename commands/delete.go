package commands

import (
	"example-bot/api"
)

func Delete(messageID string) {
	api.DeleteMessage(messageID)
}
