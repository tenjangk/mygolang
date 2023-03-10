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
func (n no) lg() {
	d := math.Log10(n.A)
	fmt.Printf("%.3f", d)
	fmt.Println("")
}

func main() {
	var (
		f, s float64
		o    string
	)

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter 1st no: ")
	fmt.Scanln(&f)
	// f, _:= reader.ReadString('\n')
	// f1, err := strconv.ParseFloat(f, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	//s1, err :=fmt.Scanln(&s)
	// s, _:= reader.ReadString('\n')
	// s1, err := strconv.ParseFloat(s, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	//fmt.Println("enter no2")

	fmt.Print(`enter the operator from +,-,*,/,^,log: `)

	fmt.Scanln(&o)
	if o == "log" {
		goto w
	} else {
		fmt.Print("enter 2nd no: ")
		fmt.Scanln(&s)
		fmt.Printf("%.2f %s %.2f = ", f, o, s)

	}
w:
	y := no{
		A: f,
		B: s,
	}

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
	case "log":
		y.lg()
	default:
		fmt.Println("invalid operator")
	}
}
