package commands

import "testing"

var testOxGameStruct = &OxGameStruct{
	Started:    false,
	MessageID:  "nankaM",
	ChannelID:  "nankaC",
	Stamps:     [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
	StampIDs:   [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
	Effects:    [][][]string{{{"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}}, {{"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}}, {{"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}, {"a", "b", "c", "d", "e"}}},
	HardMode:   false,
	FastStart:  false,
	Setsumei:   false,
	StartTime:  0,
	TimeAttack: false,
	PlayerNum:  0,
}

func TestDebug(t *testing.T) {
	Debug(testOxGameStruct)
}
