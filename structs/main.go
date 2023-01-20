package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	jangk := User{"tenzing jangk", "tj@gmail.com", 17}
	fmt.Println(jangk)
	fmt.Printf("jangk details are %+v\n", jangk)
	fmt.Printf("name: %v email: %v", jangk.Name, jangk.Email)

}

type User struct {
	Name  string
	Email string
	Age   int
}
