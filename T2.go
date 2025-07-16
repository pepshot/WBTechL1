package main

import "fmt"

func square(val int, ch chan int) {
	ch <- val * val
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	for _, num := range nums {
		go square(num, ch)
	}

	for range nums {
		fmt.Println(<-ch)
	}
}
