package main

import "fmt"

func generator(nums []int, out chan<- int) {
	defer close(out)
	for _, x := range nums {
		out <- x
	}
}

func doubler(in <-chan int, out chan<- int) {
	defer close(out)
	for x := range in {
		out <- x * 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(numbers, ch1)
	go doubler(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}
