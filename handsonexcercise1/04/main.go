package main

import "fmt"

type urtype int

var x urtype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Println(x)
}
