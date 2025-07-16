package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker №%d: %d\n", id, job)
	}
	fmt.Printf("Worker №%d exiting\n", id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	done := make(chan struct{})

	go func() {
		id := 1
		for {
			select {
			case <-done:
				close(jobs)
				return
			default:
				jobs <- id
				id++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-stop
	fmt.Println("\nReceived interrupt. Shutting down...")

	close(done)

	wg.Wait()
	fmt.Println("All workers stopped.")
}
