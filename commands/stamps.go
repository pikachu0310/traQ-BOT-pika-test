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

const stampsUsageText = "UserNameが発言したメッセージの中で、:Stamp:を付けた**人数**がMinStampNumより多いものを全部返します。(個数ではなく人数！)\nusage: @BOT_pika_test /stamps :@UserName: :Stamp: (MinStampNum)"
const defaultStampName = "w"
const defaultMinStampNum = 1

type messageWithCount = lo.Tuple2[*traq.Message, int]

// Stamps UserNameが発言したメッセージの中で、 :Stamp:を付けた人数がMinStampNumより多いものを、
// 多い順に全部コードブロックの中に返す。 top5を最後にリンクだけ出力する。
func Stamps(cmdText string, channelID string) error {
	post := func(content string) error {
		_, err := api.PostMessageWithErr(channelID, content)
		return err
	}
	userName, stampName, minStampNum, err := stampsParseArgs(cmdText)
	if err != nil {
		return post(fmt.Sprintf("%s\n%s", err.Error(), stampsUsageText))
	}
	user, err := api.GetUserByUserName(userName)
	if err != nil {
		return post(fmt.Sprintf("%s\n%s", err.Error(), stampsUsageText))
	}
	stamp, err := api.GetStampByStampName(stampName)
	if err != nil {
		return post(fmt.Sprintf("%s\n%s", err.Error(), stampsUsageText))
	}
	resultMessage, err := api.PostMessageWithErr(channelID, fmt.Sprintf("Searching...(%d):loading:", 0))
	if err != nil {
		return post(fmt.Sprintf("Failed to post message: %s", err.Error()))
	}

	// UserNameが発言したメッセージを100件ずつ取得する。
	messages, err := getUserMessages(user.Id, resultMessage.Id)
	if err != nil {
		return post(fmt.Sprintf("Message search failed: %s", err.Error()))
	}
	// 取得したそれぞれのメッセージに対して、:Stamp:を付けた人数を数える。
	messagesWithCount := countStampNumbers(messages, stamp.Id)
	// 上で数えた人数がMinStampNumより多いメッセージだけを集める。
	messagesWithCount = lo.Filter(messagesWithCount, func(m messageWithCount, _ int) bool { return m.B >= minStampNum })
	// 上で集めたメッセージを、人数の多い順に、Sortする。
	slices.SortFunc(messagesWithCount, func(a, b messageWithCount) bool { return a.B > b.B })
	// 上で集めたメッセージの数とURLと時間を、フォーマットして文字列として集める。
	formattedMessages := lo.Map(messagesWithCount, func(m messageWithCount, _ int) string {
		return fmt.Sprintf("%d https://q.trap.jp/messages/%s %s", m.B, m.A.Id, m.A.CreatedAt.Format(time.DateTime))
	})
	// 上で集めた文字列に加えて、最後に上で集めたメッセージの人数の多さTOP5のURLを追加する。
	topFiveMessages := First(messagesWithCount, 5)
	formattedTopFiveMessages := lo.Map(topFiveMessages, func(m messageWithCount, _ int) string {
		return fmt.Sprintf("https://q.trap.jp/messages/%s", m.A.Id)
	})

	resultContent := formatPostContent(formattedMessages, formattedTopFiveMessages, user.Name, stamp.Name, minStampNum)
	if err = api.EditMessageWithErr(resultMessage.Id, resultContent); err != nil {
		return post(fmt.Sprintf("Message edit failed: %s", err.Error()))
	}
	return nil
}

func stampsParseArgs(cmdText string) (userName string, stampName string, minStampNum int, err error) {
	stampName = defaultStampName
	minStampNum = defaultMinStampNum

	args := CmdArgs(cmdText)
	if len(args) == 0 {
		return "", "", 0, errors.New("args is nil")
	}
	userName = args[0]
	if strings.HasPrefix(userName, ":@") && strings.HasSuffix(userName, ":") && len(userName) >= 3 {
		userName = userName[2 : len(userName)-1]
	} else if strings.HasPrefix(userName, ":") && strings.HasSuffix(userName, ":") && len(userName) >= 2 {
		userName = userName[1 : len(userName)-1]
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
	minStampNum, err = strconv.Atoi(args[2])
	return
}

// getUserMessages UserNameがtraQ上で発言した全てのメッセージを返す
func getUserMessages(userID string, progressMessageID string) ([]*traq.Message, error) {
	var messages []*traq.Message
	var before = time.Now()
	for {
		t1 := time.Now()
		res, err := api.GetMessagesFromUser(userID, 100, 0, before)
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

// countStampNumbers 取得したそれぞれのメッセージに対して、:Stamp:を付けた人数を数える。
func countStampNumbers(messages []*traq.Message, stampID string) []messageWithCount {
	return lo.Map(messages, func(message *traq.Message, index int) messageWithCount {
		matchedStamps := lo.Filter(message.Stamps, func(m traq.MessageStamp, index int) bool {
			return m.StampId == stampID
		})
		return lo.T2(message, len(matchedStamps))
	})
}

func formatPostContent(all []string, top []string, userName string, stampName string, minStampNum int) string {
	var lines []string
	lines = append(lines, fmt.Sprintf(":%s:を付けた人数が%dより多い:@%s:が発言したメッセージの数: %d\n", stampName, minStampNum, userName, len(all)))
	lines = append(lines, "```")
	lines = append(lines, First(all, 100)...)
	lines = append(lines, "```")
	lines = append(lines, top...)
	return strings.Join(lines, "\n")
}

func First[T any](slice []T, num int) []T {
	if num >= len(slice) {
		return slice
	} else {
		return slice[:num]
	}
}
