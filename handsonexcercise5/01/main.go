package main

import "fmt"

type person struct {
	first     string
	last      string
	favFlavor []string
}

func main() {
	p1 := person{
		first:     "tem",
		last:      "mol",
		favFlavor: []string{"chocochip", "oreo"},
	}

	p2 := person{
		first:     "jan",
		last:      "tenz",
		favFlavor: []string{"blueberry", "strawberry"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavor {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavor {
		fmt.Println("\t", i, v)
	}
}
