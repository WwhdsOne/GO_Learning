package main

import (
	"encoding/json"
	"fmt"
)

type stu struct {
	Id      int    `json:"id"`
	StuName string `json:"name"`
	Age     int    `json:"age"`
	Live    bool   `json:"bool"`
}

func main() {
	var Students = [3]stu{
		{1, "wwh", 18, true},
		{2, "wwh", 18, true},
		{3, "wwh", 18, true},
	}
	fmt.Println("stu : ", Students)
	data, _ := json.Marshal(Students)
	fmt.Printf("%s\n", data)
	var decode []stu
	err := json.Unmarshal(data, &decode)
	if err != nil {
		return
	}
	fmt.Println("decode : ", decode)
}
