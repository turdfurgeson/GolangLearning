package main

import "fmt"

func main() {
	var name interface{} = "Bertney"
	//asserting that it is a string
	str, _ := name.(string)
	fmt.Printf("%T\n", str)

	var val interface{} = 42
	fmt.Printf("%T\n", val)
	//asserting that it is an int so it can be mainpulated
	fmt.Println(val.(int) + 58)
	fmt.Printf("%T\n", val)
}
