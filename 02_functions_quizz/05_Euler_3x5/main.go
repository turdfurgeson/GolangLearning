package main

import "fmt"


func nums () {

	low := 0

	fmt.Println("Enter a low whole number")

	fmt.Scan(&low)

	high := 0

	fmt.Println("Enter a high whole number")

	fmt.Scan(&high)

	num := []int{}

	for i := low; i <= high; i++ {

		if int(i)%5 == 0 || int(i)%3 == 0 {

			num = append(num, i)
		}
	}

	fmt.Println("The numbers divisible by 3 or 5 are", num)

}

func main () {

	nums()
}




/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000. */