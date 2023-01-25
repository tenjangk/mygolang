package main

import "fmt"

func main() {                   //print every no from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\t\t%b\n", i, i)
	}
}
