package main

import (
	"fmt"
	"strconv"
)

func main () {
	w := "42"
	x, _ := strconv.Atoi(w)
	fmt.Println(x)

	y := 42
	z := strconv.Itoa(y)
	fmt.Println(z)
}

//Atoi is short for converting ASCII to int. Takes string input and converts
// it to int