package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type coloredSquare struct {
	color string
	s     square
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (cs coloredSquare) area() float64 {
	return cs.s.area() * 2
}

func main() {
	c := circle{radius: 5}
	s := square{side: 4}
	cs := coloredSquare{s: square{side: 4}, color: "red"}
	fmt.Println(c.area())
	fmt.Println(s.area())
	fmt.Println(cs.area())
}
