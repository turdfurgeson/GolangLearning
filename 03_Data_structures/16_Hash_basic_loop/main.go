package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)	//This prints
		// out the integer, the letter representing it and then assigns
		// it a number of 1-12. Later we will add each letter into
		// one of 12 buckets
	}
}
