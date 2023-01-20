package main

import "fmt"

func main() {
	fmt.Println("main func")
	hello()
	hi()

	result, msg := add(3, 20)
	fmt.Println("result is ", result, msg)

	proresult := proadd(3, 4, 5, 2, 6)
	fmt.Println(proresult)

}

func proadd(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func add(first int, sec int) (int, string) {
	return first + sec, "hi from add func"
}

func hi() {
	fmt.Println("hi another method")
}

func hello() {
	fmt.Println("hellooooo go")
}
