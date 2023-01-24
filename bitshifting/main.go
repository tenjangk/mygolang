package main

import "fmt"

const (
	_  = iota             //throw away first iota i.e 0
	kb = 1 << (iota * 10) // 1 will shift over 1*10 = 10 times
	mb = 1 << (iota * 10) // 1 will shift over 2*10 = 20 times
	gb = 1 << (iota * 10) // 1 will shift over 3*10 = 30 times
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Println("")

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
