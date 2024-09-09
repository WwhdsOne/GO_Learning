package main

import "fmt"

type Circle struct {
	radius float64
}

func (circle Circle) getArea() float64 {
	return 3.14 * circle.radius * circle.radius
}

func main() {
	c := Circle{radius: 10}
	fmt.Println(c.getArea())
}
