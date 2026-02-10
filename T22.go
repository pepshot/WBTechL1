package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1048577) // 2^20 + 1
	b := big.NewInt(2097153) // 2^21 + 1

	sum := new(big.Int) // Вывод результатов
	diff := new(big.Int)
	prod := new(big.Int)
	quot := new(big.Int)

	sum.Add(a, b)

	diff.Sub(a, b)

	prod.Mul(a, b)

	if b.Cmp(big.NewInt(0)) != 0 {
		quot.Div(a, b)
	} else {
		fmt.Println("Деление на ноль невозможно")
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", prod)
	fmt.Println("a / b =", quot)
}
