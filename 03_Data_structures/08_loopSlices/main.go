package main

import "fmt"

func main() {

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)	//outer loop creates a slice
		// of int. SO total of 3 slices inside!
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)	//inner loop
			// appends whatever number 'j' is to each slice of
			// the outer loop
		}
		transactions = append(transactions, transaction)
		//then transactions is appended to add the three slices of
		// out loop that are each indexed with the numbers from the
		// inner loop
	}
	fmt.Println(transactions)
}
