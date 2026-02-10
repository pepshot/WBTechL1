package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End after 2 seconds")
}
