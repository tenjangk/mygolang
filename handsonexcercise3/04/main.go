package main

import "fmt"

func main() {
	i := 1995  // forloop using this syntax - for{}
	for {
		if i <= 2023 {
			fmt.Printf("%d\n", i)
		}
		i++
	}
}
