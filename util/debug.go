package util

import (
	"reflect"
)

type OxGameStruct struct {
	Started   bool
	MessageID string
	ChannelID string
	Stamps    [][]string
	StampIDs  [][]string
	Effects   [][][]string
	HardMode  bool
	FastStart bool
	Setsumei  bool
	StartTime int64
}

func DebugList(OxGamePlayingList []OxGameStruct) {
	message := ""
	for i := 0; i < len(OxGamePlayingList); i++ {
		OxGameDebug := OxGamePlayingList[i]
		OxGameDebugType := reflect.TypeOf(OxGameDebug)
		OxGameDebugValue := reflect.ValueOf(OxGameDebug)
		message += "```\n"
		for j := 0; j < OxGameDebugType.NumField(); j++ {
			field := OxGameDebugType.Field(j)
			value := OxGameDebugValue.Field(j)
			message += (field.Name + ": " + value.String() + "\n")
		}
		message += "```\n"
	}
	println(message)
}
