package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	var b = a[2:4] //设置b为a的切片

	fmt.Println("a : ", a) // 1 2 3 4 5
	fmt.Println("b : ", b) // 3 4

	var c = a[2:4:4]
	c = append(c, 233)
	fmt.Println("a : ", a) // 1 2 3 4 5
	//c容量为2，当append之后其实c已经指向了一个新的数组，所以a不变
	fmt.Println("c : ", c) // 3 4 233

	var d = a[2:4:5]
	d = append(d, 233)
	//此时d的容量是3，当我们append后，a和d的最后一位都被修改为了233
	fmt.Println("a : ", a)

	fmt.Println("d : ", d)

	d = append(d, 233)     // 此时d超出了最大容量指向了新的数组
	fmt.Println("a : ", a) // 1 2 3 4 233

	fmt.Println("d : ", d) // 3 4 233 233

}
