package main

import "fmt"

func main() {
	x := 16
	y := 16.023
	//converts int to float64.Widening conversion - going from whole
	// number to number with decimals. "widening" number
	fmt.Println(y * float64(x))

	//Narrowing conversion here since taking float to whole number.
	fmt.Println(x * int(y))
}
