package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46}
	b := []int{47, 48, 49, 50, 51}
	fmt.Println(a)
	fmt.Println(b)
	c := append(a[2:], b[:2]...)
	fmt.Println(c)
	d := append(a[1:], b[:1]...)
	fmt.Println(d)
}
