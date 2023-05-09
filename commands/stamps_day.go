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

const stampsDayUsageText = "UserNameが発言したメッセージの中で、:Stamp:を付けた**人数**がMinStampNumより多いものを全部返します。(個数ではなく人数！)\nusage: @BOT_pika_test /stamps :@UserName: :Stamp: (MinStampNum)"

func StampsDay(cmdText string, channelID string) error {
	post := func(content string) error {
		_, err := api.PostMessageWithErr(channelID, content)
		return err
	}
	stampName, after, before, err := stampsDayParseArgs(cmdText)
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
	topFiveMessages := First(messagesWithCount, 5)
	formattedTopFiveMessages := lo.Map(topFiveMessages, func(m messageWithCount, _ int) string {
		return fmt.Sprintf("https://q.trap.jp/messages/%s", m.A.Id)
	})

	resultContent := stampsDayFormatPostContent(formattedTopFiveMessages)
	if err = api.EditMessageWithErr(resultMessage.Id, resultContent); err != nil {
		return post(fmt.Sprintf("Message edit failed: %s", err.Error()))
	}
	return nil
}

func stampsDayFormatPostContent(messages []string) string{} {
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
		before = messages[len(messages)-1].CreatedAt
		api.EditMessage(progressMessageID, fmt.Sprintf("Searching...(%d):loading:", len(messages)))
	}

	return messages, nil
}

func stampsDayParseArgs(cmdText string) (stampName string, after time.Time, before time.Time, err error) {
	stampName = defaultStampName
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	after = midnight
	before = midnight.AddDate(0, 0, 1)

	args := CmdArgs(cmdText)
	if len(args) == 0 {
		return "", time.Time{}, time.Time{}, errors.New("args is nil")
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
		return
	}
	if len(args) == 3 {
		return
	}
	before, err = time.Parse("2006-01-02", args[3])
	if err != nil {
		return
	}
	return
}
