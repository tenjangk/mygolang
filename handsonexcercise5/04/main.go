package main

import "fmt"

func main() {
	p := struct {
		vehicle string
		no      map[string]int
		color   []string
	}{
		vehicle: "sedan",
		no: map[string]int{
			"del": 0x319,
			"pun": 0x231,
		},
		color: []string{
			"silver",
			"blue",
			"jetblack",
		},
	}
	fmt.Println(p)
	fmt.Println(p.color)
	fmt.Println(p.no)
	fmt.Println("------------")

	for k, v := range p.no {
		fmt.Println(k, v)
	}
	fmt.Println("------------")
	for i, val := range p.color {
		fmt.Println(i, val)
	}
}
