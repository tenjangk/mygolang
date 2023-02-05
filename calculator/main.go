package main

import (
	"fmt"
)

type no struct {
	A float64
	B float64
}

func (n no) sum() float64 {
	return n.A + n.B
}
func (n no) diff() float64 {
	return n.A - n.B
}
func (n no) prd() float64 {
	return n.A * n.B
}
func (n no) div() {
	if n.B == 0 {
		fmt.Println("cant divide by zero")
	} else {
		fmt.Println(n.A / n.B)
	}
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
	var f float64
	var s float64
	var operator rune
	fmt.Println("enter no1")
	fmt.Scan(f)
	fmt.Println("enter no2")
	fmt.Scan(s)
	fmt.Println(`enter the operator from 1-4 :
        1.add
	    2.diff
	    3.product
	    4.div`)
	fmt.Scan(operator)
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

	switch operator {
	case '+':
		y.sum()

	case '-':
		y.diff()

	case '*':
		y.prd()
	case '/':
		y.div()
	default:
		fmt.Println("invalid operator")
	}
}
