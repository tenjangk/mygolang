package main

import "fmt"

const (
	a = iota //iota begins with 0
	b = iota + 1
	c = iota
	z
)

const (
	d = iota + 3 //iota increments by 1
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
