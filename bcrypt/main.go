package main

import "fmt"

func main() {
	s := `password987`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginpasswd := `password987`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpasswd))
	if err != nil {
		fmt.Println("login failed")
		return
	}
	fmt.Println("u r logged in")
}
