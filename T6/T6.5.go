package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker5() {
	defer fmt.Println("defer выполнится")
	fmt.Println("Работаю...")
	time.Sleep(time.Second)
	runtime.Goexit()
	fmt.Println("Этот fmt не успеет выполниться")
}

func main() {
	go worker5()
	time.Sleep(2 * time.Second)
	fmt.Println("main завершён")
}
