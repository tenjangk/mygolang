package main

import "fmt"

const (
	a int  = 6
	b bool = true
	c      = "me"
	d      = 24.421
)

func main() {
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", b, b)
	fmt.Printf("%v\t\t%T\n", c, c)
	fmt.Printf("%v\t\t%T\n", d, d)
}
