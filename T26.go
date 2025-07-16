package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "aabcd"

	inputString = strings.ToLower(inputString)

	tableRunes := map[rune]bool{}
	for _, v := range inputString {
		_, ok := tableRunes[v]
		if ok {
			fmt.Println(false)
			return
		}
		tableRunes[v] = true
	}
	fmt.Println(true)
}
