package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("list of all languages:", languages)
	fmt.Println("js shorts for", languages["js"])

	delete(languages, "js")
	fmt.Println("list of all languages:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("for key v,value is %v \n", value)
	}

}
