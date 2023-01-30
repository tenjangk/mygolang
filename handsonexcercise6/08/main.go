package main

import "fmt"

func main() {
	a := foo()  
	fmt.Printf("%T\n", a)
	//fmt.Println(a)
	fmt.Println(a())

	c := a()
	fmt.Printf("%T\n", c)
	fmt.Println(c)
}

func foo() func() int {  //Create a func which returns a func and assign the returned func to a variable
	return func() int {
		return 24
	}

}
