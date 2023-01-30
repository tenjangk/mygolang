package main

import "fmt"

func main() {
	x := increment() //create a func which “encloses” the scope of a variable
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	z := increment()
	fmt.Println(z())
	fmt.Println(z())
	{
		y := 17
		fmt.Println(y)

	}

}
func increment() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
