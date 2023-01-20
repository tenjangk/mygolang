package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hey! this needs to go in file-learncode.in"

	file, err := os.Create("./mygofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length of file is ", length)
	defer file.Close()

	readFile("./mygofile.txt")
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilError(err)
	fmt.Println("Text data inthis file is \n", string(databyte))
}
