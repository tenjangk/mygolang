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
		first:     "tenz",
		last:      "jan",
		favFlavor: []string{"blueberry", "strawberry"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		//fmt.Println(v.last)
		for i, val := range v.favFlavor {
			fmt.Println("\t", i, val)
		}
	}
}
