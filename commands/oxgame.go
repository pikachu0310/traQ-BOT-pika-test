package commands

import (
	"example-bot/api"
	"fmt"
	"github.com/traPtitech/traq-ws-bot/payload"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type OxGameStruct struct {
	Started    bool
	MessageID  string
	ChannelID  string
	Stamps     [][]string
	StampIDs   [][]string
	Effects    [][][]string
	HardMode   bool
	FastStart  bool
	Setsumei   bool
	StartTime  int64
	TimeAttack bool
	PlayerNum  int
}

var OxGameStartStampNormal string = "type_normal"
var OxGameStartStampHard string = "crying-hard"
var OxGamePlayingList []*OxGameStruct
var OxGameVersion string = "1.0.6"

// var Effect = []string{"ex-large", "large", "small", "rotate", "rotate-inv", "wiggle", "parrot", "zoom", "inversion", "turn", "turn-v", "happa", "pyon", "flashy", "pull", "atsumori", "stretch", "stretch-v", "conga", "marquee", "conga-inv", "marquee-inv", "attract", "ascension", "shake", "party", "rainbow"}
// var Effect1 = []string{"ex-large", "large", "small"}
var Effect2 = []string{"rotate", "rotate-inv", "wiggle", "parrot", "zoom", "inversion", "turn", "turn-v", "happa", "pyon", "flashy", "pull", "atsumori", "stretch", "stretch-v", "conga", "marquee", "conga-inv", "marquee-inv", "attract", "ascension", "shake", "party", "rainbow"}

func OxGameDebug(OxGame *OxGameStruct) {
	fmt.Println("OxGameDebug")
	fmt.Println(OxGame.Started)
	fmt.Println(OxGame.MessageID)
	fmt.Println(OxGame.ChannelID)
	fmt.Println(OxGame.Stamps)
	fmt.Println(OxGame.StampIDs)
	fmt.Println(OxGame.Effects)
	fmt.Println(OxGame.HardMode)
	fmt.Println(OxGame.FastStart)
}

func OxGameStart(ChannelID string, slice []string) {
	OxGame := OxGameGet(ChannelID)
	if len(slice) == 1 {
		message := "## :blob_speedy_roll_inverse::blob_speedy_roll_inverse::blob_speedy_roll_inverse:早押しスタンプ:o::x:ゲーム Ver" + OxGameVersion + ":blob_speedy_roll::blob_speedy_roll::blob_speedy_roll:\n``@BOT_pika_test /game`` と入力することで遊べるよ！\n```\n遊び方 : BOTが3x3のマス上全てにランダムなスタンプを配置するので、\nマスと同じスタンプを押してマスを獲得し、一列揃えたら勝ち！(誰も揃わなかったら最も多かった人からランダム)\n```\n\n#### このメッセージに:type_normal:を押すとノーマルモード\n#### このメッセージに:crying-hard:を押すとハードモードで始まるよ！\n全9マスを埋めるTA(TimeAttack)モードもあるぞ！(↓のコマンドで出来る)(通常時でも全マスが埋まってたらTAモード扱いになる)\ntips:``/game start``,``/game start hard``,``/game ta``,``/game ta hard``でクイックスタート(この文章をスキップ)ができるよ！\nタイムが出るのでタイムアタックとしても楽しんで！ Enjoy! :party_blob:"
		OxGame.Setsumei = true
		OxGame.MessageID = api.PostMessage(ChannelID, message).Id
		return
	}
	if slice[1] == "start" {
		if !OxGame.Started {
			OxGame.TimeAttack = false
			if len(slice) >= 3 {
				if slice[2] == "hard" {
					OxGameMakeEffect(OxGame)
					OxGame.HardMode = true
				}
			}
			OxGame.FastStart = true
			OxGameInit(OxGame, ChannelID)
		} else {
			api.PostMessage(ChannelID, "ゲームはすでに開始されています。")
		}
	} else if slice[1] == "reset" {
		OxGameInitCompletely(OxGame)
		api.PostMessage(ChannelID, "ゲームをリセットしました。")
	} else if slice[1] == "debug" {
		OxGameDebug(OxGame)
		api.PostMessage(ChannelID, Debug(OxGame))
	} else if slice[1] == "timeattack" || slice[1] == "ta" {
		if !OxGame.Started {
			OxGame.TimeAttack = true
			if len(slice) >= 3 {
				if slice[2] == "hard" {
					OxGameMakeEffect(OxGame)
					OxGame.HardMode = true
				}
			}
			OxGame.FastStart = true
			OxGameInit(OxGame, ChannelID)
		} else {
			api.PostMessage(ChannelID, "ゲームはすでに開始されています。")
		}
	}
}

func OxGameInit(OxGame *OxGameStruct, ChannelID string) {
	OxGame.Started = true
	OxGame.ChannelID = ChannelID
	OxGame.Stamps = make([][]string, 3, 3)
	OxGame.Setsumei = false
	for i := 0; i < 3; i++ {
		OxGame.Stamps[i] = make([]string, 3, 3)
	}
	OxGame.StampIDs = make([][]string, 3, 3)
	for i := 0; i < 3; i++ {
		OxGame.StampIDs[i] = make([]string, 3, 3)
	}
	OxGameMakeStamps(OxGame)
	if OxGame.HardMode {
		OxGameFirstMessageHard(OxGame)
		nowUTC := time.Now().UTC()
		OxGame.StartTime = nowUTC.UnixNano() / int64(time.Millisecond)
	} else {
		OxGameFirstMessage(OxGame)
		nowUTC := time.Now().UTC()
		OxGame.StartTime = nowUTC.UnixNano() / int64(time.Millisecond)
	}
}

func OxGameInitCompletely(OxGame *OxGameStruct) {
	OxGame.Started = false
	OxGame.MessageID = ""
	OxGame.Stamps = make([][]string, 3, 3)
	for i := 0; i < 3; i++ {
		OxGame.Stamps[i] = make([]string, 3, 3)
	}
	OxGame.StampIDs = make([][]string, 3, 3)
	for i := 0; i < 3; i++ {
		OxGame.StampIDs[i] = make([]string, 3, 3)
	}
	OxGame.HardMode = false
	OxGame.FastStart = false
	OxGame.Setsumei = false
	OxGame.TimeAttack = false
	OxGame.StartTime = 0
}

func OxGameNew(ChannelID string) *OxGameStruct {
	newOxGame := &OxGameStruct{ChannelID: ChannelID, Started: false}
	OxGamePlayingList = append(OxGamePlayingList, newOxGame)
	return newOxGame
}

func OxGameGet(ChannelID string) *OxGameStruct {
	for i := 0; i < len(OxGamePlayingList); i++ {
		if OxGamePlayingList[i].ChannelID == ChannelID {
			return OxGamePlayingList[i]
		}
	}
	return OxGameNew(ChannelID)
}

func OxGameMakeStamps(OxGame *OxGameStruct) {
	//OxGameDebug(OxGame)
	stamps := api.GetAllStamps()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rand.Seed(time.Now().UnixNano())
			stamp := stamps[rand.Intn(len(stamps))]
			OxGame.Stamps[i][j] = stamp.Name
			OxGame.StampIDs[i][j] = stamp.Id
		}
	}
}

func OxGameFirstMessage(OxGame *OxGameStruct) {
	if !OxGame.FastStart {
		OxGameEditMessage(OxGame)
		return
	}
	message := ":" + OxGame.Stamps[0][0] + ": :" + OxGame.Stamps[0][1] + ": :" + OxGame.Stamps[0][2] + ":\n" +
		":" + OxGame.Stamps[1][0] + ": :" + OxGame.Stamps[1][1] + ": :" + OxGame.Stamps[1][2] + ":\n" +
		":" + OxGame.Stamps[2][0] + ": :" + OxGame.Stamps[2][1] + ": :" + OxGame.Stamps[2][2] + ":"
	OxGame.MessageID = api.PostMessage(OxGame.ChannelID, message).Id
}

func OxGameEditMessage(OxGame *OxGameStruct) {
	message := ":" + OxGame.Stamps[0][0] + ": :" + OxGame.Stamps[0][1] + ": :" + OxGame.Stamps[0][2] + ":\n" +
		":" + OxGame.Stamps[1][0] + ": :" + OxGame.Stamps[1][1] + ": :" + OxGame.Stamps[1][2] + ":\n" +
		":" + OxGame.Stamps[2][0] + ": :" + OxGame.Stamps[2][1] + ": :" + OxGame.Stamps[2][2] + ":"
	api.EditMessage(OxGame.MessageID, message)
}

func OxGameFirstMessageHard(OxGame *OxGameStruct) {
	message := ""
	if !OxGame.FastStart {
		OxGameEditMessageHard(OxGame)
		return
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			message += ":" + OxGame.Stamps[i][j] + "." + OxGame.Effects[i][j][0] + "." + OxGame.Effects[i][j][1] + "." + OxGame.Effects[i][j][2] + "." + OxGame.Effects[i][j][3] + "." + OxGame.Effects[i][j][4] + ":"
		}
		message += "\n"
	}
	OxGame.MessageID = api.PostMessage(OxGame.ChannelID, message).Id
}

func OxGameEditMessageHard(OxGame *OxGameStruct) {
	message := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			message += ":" + OxGame.Stamps[i][j] + "." + OxGame.Effects[i][j][0] + "." + OxGame.Effects[i][j][1] + "." + OxGame.Effects[i][j][2] + "." + OxGame.Effects[i][j][3] + "." + OxGame.Effects[i][j][4] + ":"
		}
		message += "\n"
	}
	api.EditMessage(OxGame.MessageID, message)
}

func OxGameJudge(OxGame *OxGameStruct) {
	if !OxGame.TimeAttack {
		for i := 0; i < 3; i++ {
			if OxGame.Stamps[i][0] == OxGame.Stamps[i][1] && OxGame.Stamps[i][1] == OxGame.Stamps[i][2] {
				OxGameWin(OxGame.Stamps[i][0], OxGame)
				return
			}
			if OxGame.Stamps[0][i] == OxGame.Stamps[1][i] && OxGame.Stamps[1][i] == OxGame.Stamps[2][i] {
				OxGameWin(OxGame.Stamps[0][i], OxGame)
				return
			}
		}
		if OxGame.Stamps[0][0] == OxGame.Stamps[1][1] && OxGame.Stamps[1][1] == OxGame.Stamps[2][2] {
			OxGameWin(OxGame.Stamps[0][0], OxGame)
			return
		}
		if OxGame.Stamps[0][2] == OxGame.Stamps[1][1] && OxGame.Stamps[1][1] == OxGame.Stamps[2][0] {
			OxGameWin(OxGame.Stamps[0][2], OxGame)
			return
		}
	}
	if OxGame.StampIDs[0][0] == "Done" && OxGame.StampIDs[0][1] == "Done" && OxGame.StampIDs[0][2] == "Done" &&
		OxGame.StampIDs[1][0] == "Done" && OxGame.StampIDs[1][1] == "Done" && OxGame.StampIDs[1][2] == "Done" &&
		OxGame.StampIDs[2][0] == "Done" && OxGame.StampIDs[2][1] == "Done" && OxGame.StampIDs[2][2] == "Done" {
		names := make([]string, 0, 9)
		numbers := make([]int, 0, 9)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				nameNew := true
				for k := 0; k < len(names); k++ {
					if OxGame.Stamps[i][j] == names[k] {
						numbers[k]++
						nameNew = false
					}
				}
				if nameNew {
					names = append(names, OxGame.Stamps[i][j])
					numbers = append(numbers, 1)
				}
			}
		}
		OxGame.PlayerNum = len(names)
		max := 0
		for i := 0; i < len(numbers); i++ {
			if numbers[i] == numbers[max] {
				rand.Seed(time.Now().UnixNano())
				if rand.Intn(1) == 0 {
					max = i
				}
			} else if numbers[i] > numbers[max] {
				max = i
			}
		}
		OxGame.TimeAttack = true
		OxGameWin(names[max], OxGame)
		return
	}
}

func OxGameWin(UserName string, OxGame *OxGameStruct) {
	nowUTC := time.Now().UTC()
	timeResult := (nowUTC.UnixNano() / int64(time.Millisecond)) - OxGame.StartTime
	timeSec := timeResult / 1000
	timeMilliSec := timeResult % 1000
	if OxGame.TimeAttack {
		if OxGame.HardMode {
			message := ":" + UserName + ":の勝ちです！\nタイム(HardTimeAttack): " + strconv.Itoa(int(timeSec)) + "." + strconv.Itoa(int(timeMilliSec)) + "秒でした！" + " / 人数:" + strconv.Itoa(OxGame.PlayerNum)
			api.AddMessageWithNewLine(OxGame.MessageID, message)
		} else {
			message := ":" + UserName + ":の勝ちです！\nタイム(NormalTimeAttack): " + strconv.Itoa(int(timeSec)) + "." + strconv.Itoa(int(timeMilliSec)) + "秒でした！" + " / 人数:" + strconv.Itoa(OxGame.PlayerNum)
			api.AddMessageWithNewLine(OxGame.MessageID, message)
		}
	} else {
		if OxGame.HardMode {
			message := ":" + UserName + ":の勝ちです！\nタイム(Hard): " + strconv.Itoa(int(timeSec)) + "." + strconv.Itoa(int(timeMilliSec)) + "秒でした！"
			api.AddMessageWithNewLine(OxGame.MessageID, message)
		} else {
			message := ":" + UserName + ":の勝ちです！\nタイム(Normal): " + strconv.Itoa(int(timeSec)) + "." + strconv.Itoa(int(timeMilliSec)) + "秒でした！"
			api.AddMessageWithNewLine(OxGame.MessageID, message)
		}
	}
	OxGameInitCompletely(OxGame)
}

func OxGameMakeEffect(OxGame *OxGameStruct) {
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

func OxGamePlay(MessageID string, pStamps []payload.MessageStamp) {
	OxGame := OxGameGet(api.GetMessage(MessageID).ChannelId)
	if OxGame.Setsumei {
		for k := 0; k < len(pStamps); k++ {
			stampName := api.GetStamp(pStamps[k].StampID).Name
			if stampName == OxGameStartStampNormal {
				OxGame.FastStart = false
				OxGame.Setsumei = false
				OxGameInit(OxGame, OxGame.ChannelID)
			} else if stampName == OxGameStartStampHard {
				OxGameMakeEffect(OxGame)
				OxGame.HardMode = true
				OxGame.FastStart = false
				OxGame.Setsumei = false
				OxGameInit(OxGame, OxGame.ChannelID)
			}
		}
		return
	}
	if !OxGame.Started {
		return
	}
	if MessageID != OxGame.MessageID {
		return
	}
	changed := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < len(pStamps); k++ {
				StampId := pStamps[k].StampID
				if StampId == OxGame.StampIDs[i][j] {
					OxGame.Stamps[i][j] = "@" + api.GetUser(pStamps[k].UserID).Name
					OxGame.StampIDs[i][j] = "Done"
					changed = true
				}
			}
		}
	}
	OxGameJudge(OxGame)
	if OxGame.Started && changed {
		if OxGame.HardMode {
			OxGameEditMessageHard(OxGame)
		} else {
			OxGameEditMessage(OxGame)
		}
	}
}

func Debug(OxGame *OxGameStruct) string {
	message := ""
	OxGameDebugType := reflect.TypeOf(*OxGame)
	OxGameDebugValue := reflect.ValueOf(*OxGame)
	message += "```\n"
	for i := 0; i < OxGameDebugType.NumField(); i++ {
		field := OxGameDebugType.Field(i)
		value := OxGameDebugValue.Field(i)
		message += field.Name + ": " + fmt.Sprintf("%v\n", value)
	}
	message += "```\n"
	println(message)
	return message
}
