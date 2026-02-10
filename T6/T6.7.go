package main

import (
	"fmt"
	"time"
)

func worker7() {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Горутина остановлена по time.After")
	}
}

func main() {
	go worker7()
	time.Sleep(3 * time.Second)
}
