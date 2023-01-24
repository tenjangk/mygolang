package main

import "fmt"

func main() {
	s := "hello, everyone"
	//s := `hello,
	//everyone`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // converts type string to slice of byte
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) //byte is alias for uint8

	// for i := 0; i <= len(s); i++ {
	// 	fmt.Printf("%#U", s[i])    it will print UTF8 codepoint corresponding to letter of alphabets
	// }

	for i, v := range s {
		fmt.Println(i, v) // will print index and value represented in decimal
	}

	for i, v := range s {
		fmt.Printf("at index pos %d we hav hex %#x \n", i, v)  //to print value as hexadecimal 
	}
}
