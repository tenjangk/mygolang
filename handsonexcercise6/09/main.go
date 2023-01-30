package main

import "fmt"

func main() {
	// pass a func into a func as an argument: callback
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all nos: ", sum(a...))

	b := even(sum, a...)
	fmt.Println("sum of all even nos: ", b)

}

func sum(f ...int) int {
	total := 0
	for _, v := range f {
		total += v
	}
	return total
}

func even(f func(f ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
