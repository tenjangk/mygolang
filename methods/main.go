package main

import "fmt"

func main() {
	ten := User{"tenj", "tj@gmail.com", true, 17}
	ten.GetStatus()
	ten.NewEmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)

}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("email of the user is ", u.Email)
}
