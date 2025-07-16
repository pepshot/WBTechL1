package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func work(id int, ch chan int) {
	for v := range ch {
		fmt.Printf("Worker №%d: value %d\n", id, v)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	countWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	jobs := make(chan int)

	for i := 0; i < countWorkers; i++ {
		go work(i, jobs)
	}

	id := 1
	for {
		jobs <- id
		id++
		time.Sleep(500 * time.Millisecond)
	}
}
