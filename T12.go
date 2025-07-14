package main

import "fmt"

func main() {
	words := []string{"cat", "dog", "cat", "frog", "dog", "dog"}

	set := map[string]struct{}{}
	for _, w := range words {
		_, ok := set[w]
		if !ok {
			set[w] = struct{}{}
		}
	}

	for w := range set {
		fmt.Println(w)
	}
}
