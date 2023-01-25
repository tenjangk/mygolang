package main

import "fmt"

func main() {
	if i := 10; i%5 == 0 {
		fmt.Printf("%v is a multiple of 5\n", i)
	}

	t := 22
	if t < 12 {
		fmt.Println("morning")
	} else if t == 12 {
		fmt.Println("afternoon")
	} else {
		fmt.Println("evening")
	}
}
