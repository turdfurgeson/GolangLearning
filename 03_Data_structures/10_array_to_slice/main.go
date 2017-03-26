package main

import "fmt"

func main () {

	var myArray [24]int

	mySlice := []int{}

	mySlice = append(mySlice, myArray[:])

	fmt.Println(myArray)
	fmt.Println(mySlice)
}
