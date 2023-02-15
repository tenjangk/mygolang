package main

import (
	"fmt"
	"math"
)

type no struct {
	A float64
	B float64
}

func (n no) sum() {
	fmt.Println(n.A + n.B)
}
func (n no) diff() {
	fmt.Println(n.A - n.B)
}
func (n no) prd() {
	fmt.Println(n.A * n.B)
}
func (n no) div() {
	if n.B == 0 {
		fmt.Println("cant divide by zero")
	} else {
		fmt.Println(n.A / n.B)
	}
}
func (n no) pow() {
	c := math.Pow(n.A, n.B)
	fmt.Println(c)
}

func main() {
	var (
		f, s float64
		o    string
	)
	fmt.Print("enter 1st no: ")
	_, err := fmt.Scanln(&f)
	if err != nil {
		for err != nil {
			fmt.Print("enter valid no: ")
			_, err = fmt.Scanln(&f)
		}
	}

	fmt.Print("enter 2nd no: ")
	_, errr := fmt.Scanln(&s)
	if errr != nil {
		for errr != nil {
			fmt.Print("enter valid no: ")
			_, errr = fmt.Scanln(&s)
		}
	}

	fmt.Print(`enter the operator from +,-,*,/,^: `)
	fmt.Scanln(&o)
	// f2 := float64(f1)
	// s2 := float64(s1)
	y := no{
		A: f,
		B: s,
	}
	fmt.Printf("%.2f %s %.2f = ", f, o, s)

	switch o {
	case "+":
		y.sum()

	case "-":
		y.diff()

	case "*":
		y.prd()
	case "/":
		y.div()
	case "^":
		y.pow()
	default:
		fmt.Println("invalid operator")
	}
}
