package main

import (
	"fmt"
	"sync"
)

func add(x *int, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		*x++
	}
	wg.Done()
}

func main() {
	x := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go add(&x, &wg)
	go add(&x, &wg)

	wg.Wait()
	fmt.Println(x)

	fmt.Println("Main function exiting...")
}
