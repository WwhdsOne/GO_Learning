package main

import "fmt"

const MAX = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	//指针数组
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("%d", *ptr[i])
	}
}
