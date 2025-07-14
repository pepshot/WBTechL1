package main

import (
	"fmt"
)

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	} else if len(array) == 2 {
		if array[0] < array[1] {
			return array
		}
		return []int{array[1], array[0]}
	}

	el := array[0]
	var left, middle, right []int
	for i := 0; i < len(array); i++ {
		if array[i] < el {
			left = append(left, array[i])
		} else if array[i] == el {
			middle = append(middle, array[i])
		} else {
			right = append(right, array[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	var final []int
	final = append(final, left...)
	final = append(final, middle...)
	final = append(final, right...)

	return final
}

func main() {
	array := []int{4, 2, 10, 15, 3, 10, 5, 2, 60, 90, 1000, 2, 1}
	array = quickSort(array)
	fmt.Println(array)
}
