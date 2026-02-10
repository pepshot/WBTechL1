package main

import (
	"context"
	"fmt"
	"time"
)

func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена через context:", ctx.Err())
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker3(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(time.Second)
}
