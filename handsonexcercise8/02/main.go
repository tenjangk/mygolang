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
	m := `[{"Name":"tenz","No":7834562321},{"Name":"zen","No":6834562321}]`
	bs := []byte(m)
	var u []user
	err := json.Unmarshal(bs, &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
