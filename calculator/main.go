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
	d := math.Log10(n.A) //Log10 returns the decimal logarithm of n.A
	fmt.Println(d)       //printf("%.1f") gives round up value ones place after decimal
}

//	func diff(f float64, s float64) float64 {
//		return f - s
//	}
//
//	func prd(f float64, s float64) float64 {
//		return f * s
//	}
//
//	func div(f float64, s float64) float64 {
//		if s == 0 {
//			fmt.Println("0")
//		} else {
//			return f / s
//		}
//	}
func main() {
	var (
		f float64
		s float64
		o rune
	)
	//var o rune
	//reader := bufio.NewReader(os.stdin)
	fmt.Println("enter no1")
	//f, _ := reader.ReadString('\n')
	fmt.Scan(&f)
	fmt.Println(`enter the operator from +,-,*,/,^,l`)
	//var o rune
	fmt.Scanf("%c", &o)
	if o == 'l' {
		goto w
	} else {
		fmt.Println("enter no2")
		//s, _ := reader.ReadString('\n')
		fmt.Scan(&s)
	}
w:
	y := no{
		A: f,
		B: s,
	}

	//reader := bufio.NewReader(os.Stdin)
	// fmt.Println("enter 1 no: ")
	// f,_:= reader.ReadString('\n')
	// fmt.Println("enter 2 no:")
	// s, _ := reader.ReadString('\n')
	// fmt.Println("enter operator:")
	// o, _, _ := reader.ReadRune()

	switch o {
	case '+':
		y.sum()

	case '-':
		y.diff()

	case '*':
		y.prd()
	case '/':
		y.div()
	case '^':
		y.pow()
	case 'l':
		y.lg()
	default:
		fmt.Println("invalid operator")
	}
}
