package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var channel1 = make(chan int)
	var channel2 = make(chan int)
	var done = make(chan bool)
	var sum = 0

	go func() {
		for {
			select {
			case channel1 <- rand.Int() % 10:
			case channel2 <- rand.Int() % 10:
			}
		}
	}()

	go func() {
		for {
			select {
			case num1 := <-channel1:
				sum += num1
			case num2 := <-channel2:
				sum += num2
			case <-done:
				return
			}
			fmt.Println(sum)
			if sum >= 1000 {
				done <- true
			}
		}
	}()

	time.Sleep(1 * time.Second) // 避免忙等待
}
