package main

import (
	"context"
	"fmt"
	"time"
)

func worker4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Остановка по таймауту:", ctx.Err())
			return
		default:
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker4(ctx)

	time.Sleep(3 * time.Second)
}
