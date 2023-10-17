package commands

import (
	"example-bot/api"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os/exec"
)

func Sql(ChannelID string, slice []string) {
	if len(slice) == 1 {
		return
	}
	sqlSentence := ""
	for i := 1; i < len(slice); i++ {
		sqlSentence += slice[i] + " "
	}
	out, err := exec.Command("mysql", "-t", "-N", "-u", "$NS_MARIADB_USER", "-p$NS_MARIADB_PASSWORD", "-h", "$NS_MARIADB_HOSTNAME", "$NS_MARIADB_DATABASE", "-e", sqlSentence).CombinedOutput()
	returnSentence := ""
	returnSentenceAdd := ""
	if err != nil {
		returnSentenceAdd = fmt.Sprintf("error: %s", out)
	} else {
		returnSentenceAdd = fmt.Sprintf("%s", out)
	}
	if len(returnSentenceAdd) == 0 {
		returnSentence = ":done:"
	} else {
		returnSentence = "```\n" + returnSentenceAdd + "```"
	}
	_, err = api.PostMessageWithErr(ChannelID, fmt.Sprintf(returnSentence))
	if err != nil {
		api.PostMessage(ChannelID, err.Error())
	}
}
