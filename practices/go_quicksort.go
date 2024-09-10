package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, arr[(l+r)>>1]
	for i < j {
		for {
			i++
			if arr[i] >= x {
				break
			}
		}
		for {
			j--
			if arr[j] <= x {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	go quickSort(arr, l, j)
	go quickSort(arr, j+1, r)
}
func main() {
	arr := [100000000]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(len(arr))
	}
	//fmt.Println(arr)
	start := time.Now()
	sort.Ints(arr[:])
	//fmt.Println(arr)
	speedTime := time.Since(start)
	fmt.Println(speedTime)

}
