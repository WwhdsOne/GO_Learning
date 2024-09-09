package main

import "fmt"

func add(a int) func() int {
	sum := 0
	return func() int {
		sum = sum + a
		return sum
	}
}

func main() {
	ADD := add(2)
	fmt.Println(ADD())
	fmt.Println(ADD())
}
