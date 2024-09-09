package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	a := 233
	b := 666
	fmt.Println(Max(a, b))
}
