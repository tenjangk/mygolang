package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "habi",
		Last:  "nadine",
		Age:   17,
	}
	p2 := person{
		First: "justin",
		Last:  "beiber",
		Age:   24,
	}
	a := []person{p1, p2}
	fmt.Println(a)

	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(bs))
}
