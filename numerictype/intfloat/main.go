package main

import "fmt"

// var z int
var a uint8
var b int8

func main() {
	x := 24
	y := 71.7849321
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	//z = 34.173249   it will show error since floating point value is assigned to z but z is already declared to store only type int values
	//fmt.Println(z)
	a = 255 //uint8 value is bounded from 0 to 255
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//b = -129   will show error since int8 value is bounded from -128 to 127
	b = -119
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
