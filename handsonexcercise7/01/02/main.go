package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "tenz"
	fmt.Println(*p)
	p.first = "jang"
	fmt.Println(&p.first)
	fmt.Println(*p)
}
func main() {
	p1 := person{
		first: "collen",
		last:  "hoover",
	}
	fmt.Println(p1)
	fmt.Println(&p1.first)
	changeMe(&p1)
	fmt.Println(p1)
	fmt.Println(&p1.first)
}
