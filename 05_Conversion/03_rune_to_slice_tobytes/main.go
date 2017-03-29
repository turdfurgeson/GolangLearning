package main

import "fmt"

func main() {
	// use ' ' because of the byte. Converts bytes to string
	fmt.Println(string([]byte{'h', 'i'}))
}
