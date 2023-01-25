package main

import "fmt"

func main() {
	favsports := "football"       //uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
	switch favsports {
	case "badminton":
		fmt.Println("badminton")
	case "football":
		fmt.Println("football")
	case "basketball":
		fmt.Println("basketball")
	case "swimming":
		fmt.Println("swimming")
	default:
		fmt.Println("default")

	}
}
