package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func testAtomic() {
	a := atomic.Int32{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000000; j++ {
				a.Add(1)
			}
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start).Milliseconds())
	fmt.Println(a.Load())
}

func testLock() {
	a := 0
	wg := sync.WaitGroup{}
	l := sync.Mutex{}
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000000; j++ {
				l.Lock()
				a++
				l.Unlock()
			}
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start).Milliseconds())
	fmt.Println(a)
}

func main() {
	//testAtomic()
	testLock()
}
