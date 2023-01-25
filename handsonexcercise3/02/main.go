package main

import "fmt"

func main() {                         //to print every rune code point of the upper case alphabet three times
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
