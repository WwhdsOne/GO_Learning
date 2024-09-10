package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pushNum(ch chan int) {
	ch <- rand.Intn(100)
}
func displayNum(ch <-chan int, times int) {
	for i := 0; i <= times; i++ {
		fmt.Printf("第%d个数：%d\n", i, <-ch)
	}
}
func main() {
	ch := make(chan int, 10)
	go displayNum(ch, 20)
	for i := 0; i <= 20; i++ {
		pushNum(ch)
		time.Sleep(1000 * time.Millisecond)
	}
}
