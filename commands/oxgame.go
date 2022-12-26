package commands

import (
	"example-bot/api"
	"fmt"
	"github.com/traPtitech/traq-ws-bot/payload"
	"math/rand"
	"time"
)

type OxGameStruct struct {
	Started   bool
	MessageID string
	ChannelID string
	Stamps    [][]string
	StampIDs  [][]string
	Effects   [][][]string
	HardMode  bool
}

var OxGamePlayingList []OxGameStruct

// var Effect = []string{"ex-large", "large", "small", "rotate", "rotate-inv", "wiggle", "parrot", "zoom", "inversion", "turn", "turn-v", "happa", "pyon", "flashy", "pull", "atsumori", "stretch", "stretch-v", "conga", "marquee", "conga-inv", "marquee-inv", "attract", "ascension", "shake", "party", "rainbow"}
// var Effect1 = []string{"ex-large", "large", "small"}
var Effect2 = []string{"rotate", "rotate-inv", "wiggle", "parrot", "zoom", "inversion", "turn", "turn-v", "happa", "pyon", "flashy", "pull", "atsumori", "stretch", "stretch-v", "conga", "marquee", "conga-inv", "marquee-inv", "attract", "ascension", "shake", "party", "rainbow"}

func OxGameStart(p *payload.MessageCreated, slice []string) {
	if len(slice) == 0 {
		api.PostMessage(p.Message.ChannelID, "引数が足りません")
		return
	}
	OxGame := OxGameGet(p.Message.ChannelID)
	if slice[1] == "start" {
		if len(slice) >= 2 {
			if slice[2] == "hard" {
				OxGameMakeEffect(OxGame)
				OxGame.HardMode = true
			}
		}
		if !OxGame.Started {
			OxGame.Started = true
			OxGame.ChannelID = p.Message.ChannelID
			OxGame.Stamps = make([][]string, 3, 3)
			for i := 0; i < 3; i++ {
				OxGame.Stamps[i] = make([]string, 3, 3)
			}
			OxGame.StampIDs = make([][]string, 3, 3)
			for i := 0; i < 3; i++ {
				OxGame.StampIDs[i] = make([]string, 3, 3)
			}
			api.PostMessage(p.Message.ChannelID, "ゲームを開始します。")
			OxGameMakeStamps(OxGame)
			if OxGame.HardMode {
				OxGameFirstMessageHard(OxGame)
			} else {
				OxGameFirstMessage(OxGame)
			}
		} else {
			api.PostMessage(p.Message.ChannelID, "ゲームはすでに開始されています。")
		}
	} else if slice[1] == "reset" {
		OxGame.Started = false
		api.PostMessage(p.Message.ChannelID, "ゲームをリセットしました。")
	}
}

func OxGameNew(ChannelID string) OxGameStruct {
	newOxGame := OxGameStruct{ChannelID: ChannelID, Started: false}
	OxGamePlayingList = append(OxGamePlayingList, newOxGame)
	return newOxGame
}

func OxGameGet(ChannelID string) OxGameStruct {
	for i := 0; i < len(OxGamePlayingList); i++ {
		if OxGamePlayingList[i].ChannelID == ChannelID {
			return OxGamePlayingList[i]
		}
	}
	return OxGameNew(ChannelID)
}

func OxGameMakeStamps(OxGame OxGameStruct) {
	stamps := api.GetAllStamps()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(OxGame.Stamps)
			fmt.Println(OxGame.StampIDs)
			rand.Seed(time.Now().UnixNano())
			stamp := stamps[rand.Intn(len(stamps))]
			OxGame.Stamps[i][j] = stamp.Name
			OxGame.StampIDs[i][j] = stamp.Id
		}
	}
}

func OxGameFirstMessage(OxGame OxGameStruct) {

	message := ":" + OxGame.Stamps[0][0] + ": :" + OxGame.Stamps[0][1] + ": :" + OxGame.Stamps[0][2] + ":\n" +
		":" + OxGame.Stamps[1][0] + ": :" + OxGame.Stamps[1][1] + ": :" + OxGame.Stamps[1][2] + ":\n" +
		":" + OxGame.Stamps[2][0] + ": :" + OxGame.Stamps[2][1] + ": :" + OxGame.Stamps[2][2] + ":"
	OxGame.MessageID = api.PostMessage(OxGame.ChannelID, message).Id
}

func OxGameEditMessage(OxGame OxGameStruct) {
	message := ":" + OxGame.Stamps[0][0] + ": :" + OxGame.Stamps[0][1] + ": :" + OxGame.Stamps[0][2] + ":\n" +
		":" + OxGame.Stamps[1][0] + ": :" + OxGame.Stamps[1][1] + ": :" + OxGame.Stamps[1][2] + ":\n" +
		":" + OxGame.Stamps[2][0] + ": :" + OxGame.Stamps[2][1] + ": :" + OxGame.Stamps[2][2] + ":"
	api.EditMessage(OxGame.MessageID, message)
}

func OxGameJudge(OxGame OxGameStruct) {
	for i := 0; i < 3; i++ {
		if OxGame.Stamps[i][0] == OxGame.Stamps[i][1] && OxGame.Stamps[i][1] == OxGame.Stamps[i][2] {
			api.PostMessage(OxGame.ChannelID, OxGame.Stamps[i][0]+"の勝ちです！")
			OxGame.Started = false
			return
		}
		if OxGame.Stamps[0][i] == OxGame.Stamps[1][i] && OxGame.Stamps[1][i] == OxGame.Stamps[2][i] {
			api.PostMessage(OxGame.ChannelID, OxGame.Stamps[0][i]+"の勝ちです！")
			OxGame.Started = false
			return
		}
	}
	if OxGame.Stamps[0][0] == OxGame.Stamps[1][1] && OxGame.Stamps[1][1] == OxGame.Stamps[2][2] {
		api.PostMessage(OxGame.ChannelID, OxGame.Stamps[0][0]+"の勝ちです！")
		OxGame.Started = false
		return
	}
	if OxGame.Stamps[0][2] == OxGame.Stamps[1][1] && OxGame.Stamps[1][1] == OxGame.Stamps[2][0] {
		api.PostMessage(OxGame.ChannelID, OxGame.Stamps[0][2]+"の勝ちです！")
		OxGame.Started = false
		return
	}
}

func OxGameMakeEffect(OxGame OxGameStruct) {
	OxGame.Effects = make([][][]string, 3, 3)
	for i := 0; i < 3; i++ {
		OxGame.Effects[i] = make([][]string, 3, 3)
		for j := 0; j < 3; j++ {
			OxGame.Effects[i][j] = make([]string, 5, 5)
		}
	}
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < 5; j++ {
				rand.Seed(time.Now().UnixNano())
				temp := Effect2[rand.Intn(len(Effect2))]
				for l := 0; l < len(OxGame.Effects[i][k]); l++ {
					if OxGame.Effects[i][k][j] == temp {
						temp = Effect2[rand.Intn(len(Effect2))]
						l = 0
					}
				}
				OxGame.Effects[i][k][j] = temp
			}
		}
	}
}

func OxGameFirstMessageHard(OxGame OxGameStruct) {
	message := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			message += ":" + OxGame.Stamps[i][j] + "." + OxGame.Effects[i][j][0] + "." + OxGame.Effects[i][j][1] + "." + OxGame.Effects[i][j][2] + "." + OxGame.Effects[i][j][3] + "." + OxGame.Effects[i][j][4] + ":"
		}
		message += "\n"
	}
	OxGame.MessageID = api.PostMessage(OxGame.ChannelID, message).Id
}

func OxGameEditMessageHard(OxGame OxGameStruct) {
	message := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			message += ":" + OxGame.Stamps[i][j] + "." + OxGame.Effects[i][j][0] + "." + OxGame.Effects[i][j][1] + "." + OxGame.Effects[i][j][2] + "." + OxGame.Effects[i][j][3] + "." + OxGame.Effects[i][j][4] + ":"
		}
		message += "\n"
	}
	api.EditMessage(OxGame.MessageID, message)
}

func OxGamePlay(MessageID string, pStamps []payload.MessageStamp) {
	OxGameGet(api.GetUser(MessageID).Id)
	if !OxGame.Started {
		return
	}
	if MessageID != OxGame.MessageID {
		return
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < len(pStamps); k++ {
				StampId := pStamps[k].StampID
				if StampId == OxGame.StampIDs[i][j] {
					OxGame.Stamps[i][j] = "@" + api.GetUser(pStamps[k].UserID).Name
					OxGame.StampIDs[i][j] = "done"
					if OxGame.HardMode {
						OxGameEditMessageHard(OxGame)
					} else {
						OxGameEditMessage(OxGame)
					}
				}
			}
		}
	}
	OxGameJudge(OxGame)
}