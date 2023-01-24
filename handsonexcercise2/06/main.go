package main

import "fmt"

const (
	// _ = iota + 2023                       //using iota create 4 const for next 4 yrs
	// a
	// b
	// c
	// d

	a = 2023 + iota
	b = 2023 + iota
	c = 2023 + iota
	d = 2023 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
