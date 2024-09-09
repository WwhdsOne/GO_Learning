package main

import (
	"fmt"
	"math"
)

func main() {
	sqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(sqrt(4)) // 2

}
