package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Name:%s, Age:%d", s.Name, s.Age)
}

func main() {
	s := Student{
		Name: "Tom",
		Age:  18,
	}
	fmt.Println(s)
}
