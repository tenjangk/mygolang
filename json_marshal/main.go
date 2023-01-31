package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
func main(){
	p1:=person{
		first: "habi",
		last: "nadine",
		age: 17,
	}
	p2:=person{
		first: "justin",
		last:"beiber",
		age: 24,
	}
	a:=[]person{p1,p2}
	fmt.Println(a)

	
}