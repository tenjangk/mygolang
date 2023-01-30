package main

import "fmt"

var x float64
var a = func() {    //Assign a func to a variable, then call that func
	fmt.Println("function a called from outside")
}

func main() {
	f := func() {
		fmt.Println("anonynous func called by assigning to variable f")
	}
	f()
	a()
	a = f
	a()
}
