package main

import "fmt"

func main() {
	const (
		a = iota // a = 0
		b        // b = 1
		c        // c = 2
	)
	fmt.Println(a, b, c)
	e := 233
	f := 233
	const (
		g = iota // g = 0
		h        // h = 1
	)
	fmt.Println(e, f, g, h)
}
