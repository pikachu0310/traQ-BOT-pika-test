package commands

import (
	"example-bot/util"
	"fmt"
)

func Oisu() string {
	oisu_slice := []int{0, 1, 2, 3}
	oisu_str := []string{":oisu-1:", ":oisu-2:", ":oisu-3:", ":oisu-4yoko:"}
	util.Shuffle(oisu_slice)
	var oisu string = ""
	for i := 0; i < 4; i++ {
		oisu += fmt.Sprintf(oisu_str[oisu_slice[i]])
	}
	return oisu
}
