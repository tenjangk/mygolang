package main

import "fmt"

func main() {
	s := "E"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#x\n", n)
}
