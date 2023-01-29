package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("i am ", p.first, p.last, "and ", p.age, " yrs old")
}

func main() {
	p := person{
		first: "tnz",
		last:  "jrmi",
		age:   17,
	}
	p.speak()
}
