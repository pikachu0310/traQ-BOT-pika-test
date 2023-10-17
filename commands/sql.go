package commands

import (
	"example-bot/api"
	"fmt"
	"os"
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

	// 環境変数を読み込む
	user := os.Getenv("NS_MARIADB_USER")
	password := os.Getenv("NS_MARIADB_PASSWORD")
	hostname := os.Getenv("NS_MARIADB_HOSTNAME")
	database := os.Getenv("NS_MARIADB_DATABASE")

	// ここで読み込んだ環境変数をCommandに渡す
	out, err := exec.Command("mysql", "-t", "-N", "-u", user, "-p"+password, "-h", hostname, database, "-e", sqlSentence).CombinedOutput()
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
