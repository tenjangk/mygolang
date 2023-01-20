package main

import "fmt"

func main() {

	fmt.Println("For loop in golang")

	days := []string{"sun", "mon", "tue", "wed"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("%v\n", day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		// if rougeValue == 4 {
		// 	break
		// }

		// if rougeValue == 4 {
		// 	rougeValue++
		// 	continue
		// }

		fmt.Println("Value is ", rougeValue)
		rougeValue++

	}
lco:
	fmt.Println("Jumpin at lco.in")

}
