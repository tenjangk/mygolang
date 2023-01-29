package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	defer fmt.Println(foo(xi...))

	yi := bar(xi)
	fmt.Printf("%T\t%v\n", yi, yi)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
