package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 17
	b := 24
	fmt.Println(a == b)
	fmt.Println(a <= b)
}
