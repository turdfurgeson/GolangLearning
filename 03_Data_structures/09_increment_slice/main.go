package main

import "fmt"

func main() {

	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])

	mySlice[0] = 1
	//mySlice[1] = 2 //This will yield "index out of range" error because
	// the legnth and cap are set to 1.
	fmt.Println(mySlice[0])

	mySlice[0]++
	//mySlice[0] += 1
	//mySlice[0] = mySlice[0] + 1	These two will yield same result.
	// Just different way to write code
	fmt.Println(mySlice[0])

}
