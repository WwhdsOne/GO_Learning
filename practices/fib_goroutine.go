package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fibRoutine(number int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时通知 WaitGroup 任务已完成
	x, y := 1, 1
	for i := 0; i < number; i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- x
}

func rece(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时通知 WaitGroup 任务已完成
	for i := 1; i < 15; i++ {
		fmt.Printf("第%d项：%d\n", i, <-ch)
	}
}

func main() {
	start := time.Now()
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1) // 添加一个等待任务
	go rece(ch, &wg)

	for i := 1; i < 15; i++ {
		wg.Add(1) // 添加一个等待任务
		go fibRoutine(i, ch, &wg)
	}

	wg.Wait() // 等待所有任务完成
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
