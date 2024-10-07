package main

import "fmt"

type NewString interface {
	Length() int
	DecreaseLength(num int) string
}

type normalString struct {
	str string
}

func (ns normalString) Length() int {
	return len(ns.str)
}

func (ns normalString) DecreaseLength(num int) string {
	return ns.str[:len(ns.str)-num]
}

func display(ns NewString) {
	fmt.Println(ns.Length())
	fmt.Println(ns.DecreaseLength(2))
}

func main() {
	ns := normalString{"Hello World"}
	display(ns)
}
