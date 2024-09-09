package main

import "fmt"

func main() {
	a := map[string]int{
		"LOL": 233,
		"CPP": 222,
	}
	fmt.Println(a)
	b := make(map[string]int)
	b["233"] = 233
	b["333"] = 333
	fmt.Println(b)
	delete(b, "233")
	delete(b, "333")
	fmt.Println(b)
	b["一"] = 1
	b["二"] = 2
	for zh, num := range b {
		fmt.Println(zh, num)
	}
}
