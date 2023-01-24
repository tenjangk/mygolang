package main

import "fmt"

func main() {
	a := (24 <= 17)
	b := (24 >= 17)
	c := (24 == 17)
	d := (24 != 17)
	e := (24 < 17)
	f := (24 > 17)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
