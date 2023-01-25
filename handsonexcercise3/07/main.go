package main

import "fmt"

func main() {
	switch { //without switch expression specified
	case (5%2 == 0):
		fmt.Println("even")
	case !(3%2 == 0):
		fmt.Println("prints")
		fallthrough
	case (1%2 == 0):
		fmt.Println("prints coz of fallthrough")
	default:
		fmt.Println("default")
	}
}
