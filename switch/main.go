package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and u can open")
	case 2:
		fmt.Println("Move 2 spot")
	case 3:
		fmt.Println("Move 3 spot")
		fallthrough
	case 4:
		fmt.Println("Move 4 spot")
		fallthrough
	case 5:
		fmt.Println("Move 5 spot")
	case 6:
		fmt.Println("Move 6 spot and roll dice again")
	default:
		fmt.Println("whats that")
	}
}
