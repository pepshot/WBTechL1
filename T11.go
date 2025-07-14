package main

import (
	"fmt"
	"slices"
)

func main() {
	setA := []int{1, 2, 3, 4}
	setB := []int{5, 4, 3, 8, 9}

	setIntersection := []int{}
	for i := 0; i < len(setA); i++ {
		if slices.Contains(setB, setA[i]) {
			setIntersection = append(setIntersection, setA[i])
		}
	}

	fmt.Println(setIntersection)
}
