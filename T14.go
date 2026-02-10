package main

import "fmt"

func detectType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int:
		fmt.Println("Тип переменной: chan int")
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

func main() {
	a := 42
	b := "hello"
	c := true
	d := make(chan int)

	detectType(a)
	detectType(b)
	detectType(c)
	detectType(d)
}
