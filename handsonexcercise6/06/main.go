package main

import "fmt"

func main() {
	func() {  //Build and use an anonymous func
		fmt.Println("anonymous fn ran")
	}()
	x, y := func(x int, y float64) (int, float64) {
		return x, y
	}(2, 1.7)
	fmt.Println(x, y)
}
