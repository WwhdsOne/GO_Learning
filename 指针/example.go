package main

import "fmt"

func main() {
	var a int16 = 233
	fmt.Printf("变量地址:%x\n\n", &a)
	var b = &a
	fmt.Printf("指针地址:%d\n", *b)
	var ptr *int16
	fmt.Printf("空指针:%x\n\n", ptr)
}
