package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hello, i`m %s\n", h.name)
}

func (h Human) Name() string {
	return h.name
}

func (h Human) Age() int {
	return h.age
}

type Action struct {
	Human
}

func main() {
	human := Human{"Alex", 19}
	action := Action{human}
	action.SayHi()
	fmt.Printf("Name is %s and Age is %d\n", action.Name(), action.Age())
}
