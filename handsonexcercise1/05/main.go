package main

import "fmt"

type T1 int

var x T1
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 17
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
