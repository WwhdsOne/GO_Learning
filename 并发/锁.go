package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	x := 0
	n := 100000
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			mutex.Lock()
			x++
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			mutex.Lock()
			x++
			mutex.Unlock()
		}
	}()
	go func() {
		for i := 0; i < n/10000; i++ {
			fmt.Println(x)
		}
	}()
	wg.Wait()
	fmt.Println(x)

}
