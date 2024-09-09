package main

import "fmt"

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	Hightlow(2, 0)
	fmt.Println("Program finished successfully!")
}
