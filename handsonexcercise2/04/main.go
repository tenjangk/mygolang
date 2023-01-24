package main

import "fmt"

func main() {
	a := 24
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)
	c := a << 1 //shifts the bits of that int over 1 pos to left
	fmt.Printf("%d\t\t%b\t\t%#x\n", c, c, c)
}
