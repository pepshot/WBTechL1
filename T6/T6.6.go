package main

import (
	"fmt"
)

func worker(ch <-chan int) {
	for v := range ch {
		fmt.Println("Получено:", v)
	}
	fmt.Println("Канал закрыт, горутина остановлена")
}

func main() {
	ch := make(chan int)

	go worker(ch)

	ch <- 1
	ch <- 2
	close(ch)
}
