package main

import "fmt"

func main() {
	var x rune = 'c' //rune is just an alias for int32
	var y int32 = 'd'
	fmt.Println(x, y)	//returns ASCII value of x and y
	fmt.Println(string(x))	//converts to just a string
	fmt.Println(string(y))

}
