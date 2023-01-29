package main

import "fmt"

func main() {
	f := foo(24)
	fmt.Printf("%T\t%v\n", f, f)
	b, a := bar(17, "jrmi")
	fmt.Println(b, a)
}

func foo(n int) int {
	return n
}

func bar(n int, s string) (int, string) {
	return n, s
}
