package main

import (
	"errors"
	"fmt"
)

var BigError = errors.New("number is to big")

func fib(n int) (int, error) {
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	if b > 100 {
		return b, nil
	} else {
		return b, BigError
	}
}
func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("error:", handler)
		}
	}()
	num, err := fib(13)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
