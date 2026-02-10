package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2

	fmt.Println("До удаления:", s)

	copy(s[i:], s[i+1:])
	// Обнуление последнего элемента (чтобы GC мог его собрать)
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	fmt.Println("После удаления:", s)
}
