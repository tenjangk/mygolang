package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to webrequest in golang")

	response, err := http.Get(url)
	checkError(err)

	fmt.Printf("response is of type: %T \n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkError(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
