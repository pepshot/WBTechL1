package main

import (
	"errors"
	"fmt"
)

func setBit(x int64, i uint) int64 {
	return x | (1 << i)
}

func clearBit(x int64, i uint) int64 {
	return x &^ (1 << i)
}

func main() {
	var x int64
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&x)
	if !errors.Is(err, nil) {
		return
	}

	var ind uint
	fmt.Print("Введите позицию: ")
	_, err = fmt.Scan(&ind)
	if !errors.Is(err, nil) {
		return
	}

	fmt.Printf("Исходное: %d, в довичном представлении: %b\n", x, x)

	x = clearBit(x, ind)
	fmt.Println("Очистить 0-й бит:", x)

	x = setBit(x, ind)
	fmt.Println("Установить 1-й бит:", x)
}
