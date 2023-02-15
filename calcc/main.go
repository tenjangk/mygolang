package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	var o string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter 1st no: ")

	f, _ := reader.ReadString('\n')

	f1, err := strconv.ParseFloat(strings.TrimSpace(f), 64)

	for err != nil {
		//fmt.Println("error: ", err)
		fmt.Print("enter valid no: ")
		f, _ = reader.ReadString('\n')

		f1, err = strconv.ParseFloat(strings.TrimSpace(f), 64)

	}

	//s1, err := fmt.Scanln(&s)
	fmt.Print("enter 2nd no: ")

	s, _ := reader.ReadString('\n')

	s1, err1 := strconv.ParseFloat(strings.TrimSpace(s), 64)

	for err1 != nil {
		//fmt.Println("error: ", err)
		fmt.Print("enter valid no: ")

		s, _ = reader.ReadString('\n')

		s1, err1 = strconv.ParseFloat(strings.TrimSpace(s), 64)
	}

	fmt.Print(`enter the operator from +,-,*,/,^: `)

	fmt.Scanln(&o)

	y := no{
		A: f1,
		B: s1,
	}

	fmt.Printf("%.2f %s %.2f = ", f1, o, s1)

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
