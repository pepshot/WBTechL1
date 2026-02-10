package main

import (
	"fmt"
)

func reverseRunes(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseWordsInPlace(str string) string {
	runes := []rune(str)
	n := len(runes)

	reverseRunes(runes, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	input := "snow dog sun"
	output := reverseWordsInPlace(input)
	fmt.Println(output)
}
