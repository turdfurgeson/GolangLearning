package main

import "fmt"


func main () {

	var x [58]string	//stores ASCII A

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)	//sets 'A' to index 0, and then loops
	}


	fmt.Println(x)	//prints alphabet
}