package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket 代表一个令牌桶，用于控制流量。
type TokenBucket struct {
	ch        chan struct{}
	tokens    int
	maxTokens int
	interval  time.Duration
	closeChan chan struct{}
	wg        sync.WaitGroup
}

// NewTokenBucket 创建一个新的令牌桶，指定令牌的最大数量和令牌添加的时间间隔。
func NewTokenBucket(maxTokens int, interval time.Duration) *TokenBucket {
	return &TokenBucket{
		ch:        make(chan struct{}),
		tokens:    maxTokens,
		maxTokens: maxTokens,
		interval:  interval,
		closeChan: make(chan struct{}),
	}
}

// Start 开始令牌桶的工作，它会定期向桶中添加令牌。
func (tb *TokenBucket) Start() {
	tb.wg.Add(1)
	go func() {
		defer tb.wg.Done()
		ticker := time.NewTicker(tb.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				tb.addToken()
			case <-tb.closeChan:
				return
			}
		}
	}()
}

// addToken 向令牌桶中添加一个令牌。
func (tb *TokenBucket) addToken() {
	if tb.tokens < tb.maxTokens {
		tb.tokens++
		fmt.Println("添加一个令牌")
	}
}

// Pour 尝试从令牌桶中取出一个令牌来执行任务，如果没有令牌可用，则阻塞。
// Pour 尝试从令牌桶中取出一个令牌来执行任务，如果没有令牌可用，则阻塞。
func (tb *TokenBucket) Pour() {
	tb.mu.Lock()
	if tb.tokens > 0 {
		tb.tokens-- // 取出一个令牌
		tb.mu.Unlock()
		fmt.Println("执行一个任务")
	} else {
		tb.mu.Unlock()
		fmt.Println("等待令牌...")
		<-tb.ch // 阻塞，直到有令牌可用
		fmt.Println("执行一个任务")
	}
}

// Close 停止令牌桶的工作。
func (tb *TokenBucket) Close() {
	close(tb.closeChan)
	tb.wg.Wait()
}

func main() {
	// 创建一个令牌桶，每分钟发放15个令牌。
	tb := NewTokenBucket(15, 10*time.Second) // 每秒钟发放一个令牌，总共15个令牌
	tb.Start()

	// 模拟发送13个任务
	for i := 0; i < 13; i++ {
		tb.Pour()
		fmt.Printf("任务 %d 已尝试执行\n", i+1)
	}

	// 等待一段时间，让令牌桶中的所有任务被执行
	time.Sleep(3 * time.Second)
	tb.Close()
}
