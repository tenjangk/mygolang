package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to sclice code")

	var fruitList = []string{"apple", "tomato", "peach", "banana"}

	fmt.Println("fruitlist is", fruitList)
	fmt.Printf("type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	var index = 1
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 346
	highscores[2] = 678
	highscores[3] = 236
	//highscores[4] = 000

	highscores = append(highscores, 444, 555, 777)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

}
