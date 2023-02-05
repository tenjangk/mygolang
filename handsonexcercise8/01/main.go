package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	No   int
}

func main() {
	u1 := user{
		Name: "tenz",
		No:   7834562321,
	}
	u2 := user{
		Name: "zen",
		No:   6834562321,
	}
	users := []user{u1, u2}
	m, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(m))
}
