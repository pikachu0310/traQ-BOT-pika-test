package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShuffleTest(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

var Effect2 = []string{"rotate", "rotate-inv", "wiggle", "parrot", "zoom", "inversion", "turn", "turn-v", "happa", "pyon", "flashy", "pull", "atsumori", "stretch", "stretch-v", "conga", "marquee", "conga-inv", "marquee-inv", "attract", "ascension", "shake", "party", "rainbow"}

type OxGameStruct struct {
	Started   bool
	MessageID string
	ChannelID string
	Stamps    [][]string
	StampIDs  [][]string
	Effects   [][][]string
	HardMode  bool
}

var OxGame OxGameStruct = OxGameStruct{Started: false}

func OxGameMakeEffect() {
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

func timeTest() {
	now := time.Now()
	nowUTC := now.UTC()

	fmt.Println(nowUTC.UnixNano())                           // ナノ秒
	fmt.Println(time.Millisecond)                            // 1ms
	fmt.Println(int64(time.Millisecond))                     // 1000000
	fmt.Println(nowUTC.Unix())                               // 秒
	fmt.Println(nowUTC.UnixNano() / int64(time.Millisecond)) // ミリ秒を算出する
}

func main() {

	timeTest()
}
