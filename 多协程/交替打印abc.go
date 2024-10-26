package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		for i := 0; i < n; i++ {
			<-ch1
			fmt.Print("a")
			ch2 <- true
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-ch2
			fmt.Print("b")
			ch3 <- true
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-ch3
			fmt.Print("c")
			ch1 <- true
		}
		wg.Done()
	}()

	go func() {
		ch1 <- true // 启动第一个 Goroutine

	}()
	wg.Wait()
}
