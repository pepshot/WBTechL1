package main

import (
	"fmt"
	"time"
)

func worker1(stop *bool) {
	for {
		if *stop {
			fmt.Println("Горутина остановлена по условию")
			return
		}
		fmt.Println("Работаю...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	stop := false

	go worker1(&stop)

	time.Sleep(2 * time.Second)
	stop = true

	time.Sleep(time.Second)
}
