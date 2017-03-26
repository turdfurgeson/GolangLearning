package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4}
	myOtherSlice := []int{5, 6, 7, 8}
	fmt.Println(mySlice)

	mySlice = append(mySlice, myOtherSlice...) //the ... appends the
	// entire slice to another slice

	fmt.Println(mySlice)

	mySlice = append(mySlice[:4], mySlice[7:]...)	//how to DELETE from
	// a slice. This deletes everything after index 4 until inedx 7. You
	// need the '...'
	fmt.Println(mySlice)
}
