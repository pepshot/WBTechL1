package main

import (
	"fmt"
)

var (
	array = []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21}
)

func binaryFind(value int, leftInd, rightInd int) int {
	if rightInd-leftInd == 1 {
		if value == array[leftInd] {
			return leftInd
		}
		if value == array[rightInd] {
			return rightInd
		}
		return -1
	}
	middleIndex := (rightInd - leftInd) / 2
	if array[middleIndex] == value {
		return middleIndex
	} else if value < array[middleIndex] {
		return binaryFind(value, leftInd, middleIndex-1)
	} else {
		return binaryFind(value, middleIndex+1, rightInd)
	}
}

func main() {
	index := binaryFind(7, 0, len(array)-1)
	if index == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Printf("Index value is %d\n", index)
	}
}
