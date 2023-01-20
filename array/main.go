package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println("fruitlist is", fruitList)
	fmt.Println("fruitlist is", len(fruitList))

	var vegList = [5]string{"potato", "tomato"}

	fmt.Println("veglist is", len(vegList))

}
