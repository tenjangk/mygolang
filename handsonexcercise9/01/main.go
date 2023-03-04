package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine()) //it already has 1 main goroutine

	wg.Add(2)
	go foo() //launches 2 additional goroutines
	go bar()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	wg.Wait() //make sure each goroutine finishes before your program exists
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}
