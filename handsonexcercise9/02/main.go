package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println(p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "tenz",
		age:  2,
	}
	p2 := person{
		name: "jan",
		age:  7,
	}

	saySomething(&p1) //you CAN pass a value of type *person into saySomething
	saySomething(&p2) 
	//saySomething(p2)  you CANNOT pass a value of type person into saySomething
}
