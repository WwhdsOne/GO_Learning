package main

import (
	"fmt"
	"math/rand"
)

func quickSort(ptr []int, l, r int) {
	if l >= r {
		return
	}
	var i, j, x = l, r, ptr[(l+r)>>1]
	for i <= j {
		for ptr[i] < x {
			i++
		}
		for ptr[j] > x {
			j--
		}
		if i <= j {
			ptr[i], ptr[j] = ptr[j], ptr[i]
			i++
			j--
		}
	}
	quickSort(ptr, l, i-1)
	quickSort(ptr, i, r)
}

func main() {
	// 创建一个包含20个随机数的数组
	var arr [20]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // 使用 rand.Intn 生成 0 到 99 之间的随机数
	}
	fmt.Println("Original array:", arr)

	quickSort(arr[:], 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}
