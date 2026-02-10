package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Размер map:", len(m))
}

// go run -race T7.go
