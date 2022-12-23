package main

import (
	fmt "fmt"
	"math/rand"
)

func ShuffleTest(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var test []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ShuffleTest(test)
	for i := 0; i < 10; i++ {
		fmt.Print(i, test)
	}
}
