package main

import "fmt"

func main() {

	mySlice := []string{"a", "b", "c", "d", "e", "f", "g"}

	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  //slice a slice - only returns indices 2
	// and 3
	fmt.Println(mySlice[2])    //index a slice - return index 2
	fmt.Println("myString"[2]) //index a string - returns 83, which is
	// 'S' in ASCII

}
