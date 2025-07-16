package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	timeout := 5 * time.Second

	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	stop := time.After(timeout)

	//Read from channel
	go func() {
		for val := range ch {
			fmt.Printf("Message: %d\n", val)
		}
	}()

	id := 0
	for {
		select {
		case <-stop:
			fmt.Println("Exititng")
			close(ch)
			return
		case <-ticker.C:
			ch <- id
			id++
		}
	}
}
