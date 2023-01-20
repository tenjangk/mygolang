package main

import "fmt"

func main() {
	myNumber := 24
	var pointer = &myNumber

	fmt.Println("VALUE OF acutal pOINTER IS", pointer)
	fmt.Println("VALUE OF actual pOINTER IS", *pointer)

	*pointer = *pointer - 7
	fmt.Println("new VALUE OF OINTER IS", myNumber)

}
