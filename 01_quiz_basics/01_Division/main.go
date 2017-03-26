package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Print("Please enter a number: ")
	fmt.Scan(&a)

	fmt.Print("Please enter another number: ")
	fmt.Scan(&b)

	var c = a / b

	fmt.Println(a + "divided by" + b " = " + c)

}
