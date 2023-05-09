package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slices"

	"example-bot/api"
)

const stampsDayUsageText = "指定した期間内(または今日)の全メッセージの中で:Stamp:をつけた**人数**が多いメッセージのTOP5(数は指定可能)を出力します。\n(※期間を省略すると今日で検索します。Stampを省略すると:w:で検索します。)\nusage: @BOT_pika_test /stamps day :Stamp: 2023-05-01 2023-05-07 5"

func StampsDay(cmdText string, channelID string) error {
	post := func(content string) error {
		_, err := api.PostMessageWithErr(channelID, content)
		return err
	}
	stampName, after, before, err, top := stampsDayParseArgs(cmdText)
	stamp, err := api.GetStampByStampName(stampName)
	if err != nil {
		return post(fmt.Sprintf("%s\n%s", err.Error(), stampsDayUsageText))
	}
	resultMessage, err := api.PostMessageWithErr(channelID, fmt.Sprintf("Searching...(%d):loading:", 0))
	if err != nil {
		return post(fmt.Sprintf("Failed to post message: %s", err.Error()))
	}

	// After以降Before以前に投稿されたメッセージを全て取得する。
	messages, err := getMessagesByPeriod(after, before, resultMessage.Id)
	// 取得したそれぞれのメッセージに対して、:Stamp:を付けた人数を数える。
	messagesWithCount := countStampNumbers(messages, stamp.Id)
	// 上で集めたメッセージを、人数の多い順に、Sortする。
	slices.SortFunc(messagesWithCount, func(a, b messageWithCount) bool { return a.B > b.B })
	// 上で集めた文字列に加えて、最後に上で集めたメッセージの人数の多さTOP5のURLを追加する。
	topFiveMessages := First(messagesWithCount, top)
	formattedTopFiveMessages := lo.Map(topFiveMessages, func(m messageWithCount, _ int) string {
		return fmt.Sprintf("https://q.trap.jp/messages/%s", m.A.Id)
	})

	resultContent := stampsDayFormatPostContent(formattedTopFiveMessages)
	if err = api.EditMessageWithErr(resultMessage.Id, resultContent); err != nil {
		return post(fmt.Sprintf("Message edit failed: %s", err.Error()))
	}
	return nil
}

func stampsDayFormatPostContent(top []string) string {
	var lines []string
	lines = append(lines, top...)
	return strings.Join(lines, "\n")
}

func getMessagesByPeriod(after time.Time, before time.Time, progressMessageID string) ([]*traq.Message, error) {
	var messages []*traq.Message
	var searchBefore = before
	for {
		t1 := time.Now()
		res, err := api.GetMessagesFromPeriod(after, searchBefore, 100, 0)
		fmt.Println(time.Since(t1))
		if err != nil {
			return nil, err
		}
		if len(res.Hits) == 0 {
			break
		}

		for i := range res.Hits {
			messages = append(messages, &res.Hits[i])
		}
		time.Sleep(time.Millisecond * 100)
		searchBefore = messages[len(messages)-1].CreatedAt
		api.EditMessage(progressMessageID, fmt.Sprintf("Searching...(%d):loading:", len(messages)))
	}

	return messages, nil
}

func stampsDayParseArgs(cmdText string) (stampName string, after time.Time, before time.Time, err error, top int) {
	stampName = defaultStampName
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	after = midnight
	before = after.AddDate(0, 0, 1)
	top = 5

	args := CmdArgs(cmdText)
	if len(args) == 0 {
		err = errors.New("args is nil")
		return
	}

	if len(args) == 1 {
		return
	}
	stampName = args[1]
	if strings.HasPrefix(stampName, ":") && strings.HasSuffix(stampName, ":") && len(stampName) >= 2 {
		stampName = stampName[1 : len(stampName)-1]
	}
	if len(args) == 2 {
		return
	}
	after, err = time.Parse("2006-01-02", args[2])
	if err != nil {
		before = after.AddDate(0, 0, 1)
		return
	}
	if len(args) == 3 {
		return
	}
	before, err = time.Parse("2006-01-02", args[3])
	if err != nil {
		return
	}
	if len(args) == 4 {
		return
	}
	top, err = strconv.Atoi(args[4])
	if err != nil {
		top = 5
		return
	}
	return
}
