package main

import (
	"fmt"
	"math"
)

type square struct {
	len   float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.len * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(sh shape) {
	switch sh.(type) {
	case square:
		fmt.Println("area of square: ", sh.area())
	case circle:
		fmt.Println("area of circle: ", sh.area())
	default:
		fmt.Println("neither circle non square")
	}

}

func main() {
	s1 := square{
		len:   2,
		width: 2,
	}
	c1 := circle{
		radius: 4.2,
	}
	info(c1)
	info(s1)
}
