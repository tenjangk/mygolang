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
	s := `[{"First":"T","Last":"J","Age":7},{"First":"T","Last":"y","Age":4}]`
	bs := []byte(s)
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error: ", err)
	}
	for i, v := range people {
		fmt.Println("person no: ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
	fmt.Println("all of data: ", people)
}
