package main

import (
	"fmt"
	"time"
)

func worker2(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Горутина остановлена через канал")
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})

	go worker2(done)

	time.Sleep(2 * time.Second)
	close(done)

	time.Sleep(time.Second)
}
