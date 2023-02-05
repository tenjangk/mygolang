package main

import (
	"encoding/json"
	"fmt"
	"os"
) //encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

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
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(users)
}
