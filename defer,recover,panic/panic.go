package main

import "fmt"

func Hightlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	Hightlow(high, low+1)
}

func main() {
	Hightlow(2, 0)
	fmt.Println("Program finished successfully!")
}
