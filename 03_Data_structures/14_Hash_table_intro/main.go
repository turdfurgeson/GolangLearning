package main

import "fmt"

func main() {

	letter := 'A' //single quotes denotes letter. double or backtick
	// would mean string
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)	//type is int32 bc it is utf8 which
	// is a 1-4 byte coding scheme
}
