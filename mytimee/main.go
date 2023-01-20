package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday "))

	createdDate := time.Date(2022, time.November, 24, 07, 14, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
