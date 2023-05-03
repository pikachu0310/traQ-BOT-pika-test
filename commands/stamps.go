package commands

import (
	"example-bot/api"
	"fmt"
	"github.com/traPtitech/go-traq"
	"sort"
	"strconv"
	"time"
)

func Kusa(cmdText string, channelID string) {
	var counts map[int][]string
	usageMessage := "UserNameが発言したメッセージの中で、:Stamp:を付けた**人数**がMinStampNumより多いものを全部返します。(個数ではなく人数！)\nusage: @BOT_pika_test /stamps :@UserName: :Stamp: (MinStampNum)"
	texts := StringToSlice(cmdText)
	user := &traq.User{}
	if len(texts) == 0 {
		api.PostMessage(channelID, usageMessage)
		return
	}
	user, err := api.GetUserByUserName(texts[0])
	if err != nil {
		if len(texts[0]) <= 1 {
			api.PostMessage(channelID, usageMessage)
			return
		}
		user, err = api.GetUserByUserName(texts[0][1:])
		if err != nil {
			if len(texts[0]) <= 3 {
				api.PostMessage(channelID, usageMessage)
				return
			}
			user, err = api.GetUserByUserName(texts[0][2 : len(texts[0])-1])
			if err != nil {
				api.PostMessage(channelID, err.Error()+"\n"+usageMessage)
				return
			}
		}
	}
	stampID := "6308a443-69f0-45e5-866f-56cc2c93578f"
	if len(texts) >= 2 {
		if len(texts[1]) <= 2 {
			api.PostMessage(channelID, usageMessage)
			return
		}
		stamp, err := api.GetStampByStampName(texts[1][1 : len(texts[1])-1])
		if err != nil {
			api.PostMessage(channelID, err.Error()+"\n"+usageMessage)
			return
		}
		stampID = stamp.Id
	}
	minStampNum := 1
	if len(texts) >= 3 {
		minStampNum, err = strconv.Atoi(texts[2])
		if err != nil {
			api.PostMessage(channelID, usageMessage)
			return
		}
	}
	stamp := api.GetStamp(stampID)

	postedMessage := api.PostMessage(channelID, fmt.Sprintf("Searching...(0):loading: :@%s: :%s: %d", user.Name, stamp.Name, minStampNum))
	returnMessage := ""
	var offset int32 = 0
	for {
		fmt.Println(offset)
		messages, err := api.GetMessagesFromUserNameAndLimitAndOffset(user.Id, 100, offset)
		if err != nil {
			api.EditMessage(postedMessage.Id, err.Error())
			break
		}
		if len(messages.Hits) == 0 {
			break
		}

		for _, message := range messages.Hits {
			count := 0
			for _, stamp := range message.Stamps {
				if stamp.StampId == stampID {
					count++
				}
			}
			if count >= minStampNum {
				counts[count] = append(counts[count], fmt.Sprintf("https://q.trap.jp/messages/%s", message.Id))
				returnMessage += fmt.Sprintf("%d https://q.trap.jp/messages/%s %s\n", count, message.Id, message.CreatedAt.Format("2006-01-02-15:04"))
			}
		}
		time.Sleep(time.Millisecond * 200)
		offset += 100
		if len(returnMessage) >= 9980 {
			break
		}
		api.EditMessage(postedMessage.Id, fmt.Sprintf("Searching...(%d):loading: :@%s: :%s: %d\n```\n%s```", offset, user.Name, stamp.Name, minStampNum, returnMessage))
	}
	time.Sleep(time.Millisecond * 1000)
	err = api.EditMessageWithErr(postedMessage.Id, "```\n"+returnMessage+"```")
	if err != nil {
		api.EditMessage(postedMessage.Id, "```\n"+returnMessage[:9980]+"\n......\n```")
		return
	}
}

func SortMap(m map[int][]string) []string {
	keys := getKeys(m)
	sort.Ints(keys)

	s := []string{}
	for key := range keys {
		for _, value := range m[key] {
			s = append(s, value)
		}
	}

	return s
}

func getKeys(m map[int][]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
