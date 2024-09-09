package main

import "fmt"

// 全局变量
var l = 2

func main() {
	// 变量定义
	var a int
	var b int
	// 类型推导
	var c = 1
	var d = "LOL"
	// 简短变量声明
	e := 1
	f := "LOL"
	// 多变量声明
	var g, h int
	var i, j = 1, 2
	k, l := 1, 2
	// 匿名变量
	_, m := 1, 2
	// 常量
	const n = 1
	// 输出
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
}
