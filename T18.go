package main

import (
	"fmt"
	"sync"
)

type Data struct {
	data map[string]int
	mu   sync.Mutex
}

func (d *Data) Inc(field string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[field]++
}

func main() {
	data := Data{
		data: map[string]int{
			"field": 0,
		},
		mu: sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			data.Inc("field")
		}(i)
	}

	wg.Wait()

	fmt.Println(data.data["field"])
}
