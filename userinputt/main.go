package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcom user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating:")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("type of this rating is %T", input)
}
