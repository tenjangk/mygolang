package main

import (
	"fmt"
)

type square struct {
	len   float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() {
	a := s.len * s.width
	fmt.Println("area of square: ", a)
}

func (c circle) area() {
	a := 3.14159 * (c.radius * c.radius)
	fmt.Println("area of circle: ", a)
}

type shape interface {
	area()
}

func info(sh shape) {
	fmt.Println(sh)
}

func main() {
	s1 := square{
		len:   2,
		width: 2,
	}
	c1 := circle{
		radius: 4.2,
	}
	info(s1)
	info(c1)
}
