package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if i%5 == 0 && i%3 == 0 {

			fmt.Println("fizzbuzz")

		} else if i%5 == 0 {

			fmt.Println("buzz")

		} else if i%3 == 0 {

			fmt.Println("fizz")

		} else {

			fmt.Println(i)

		}
	}
}
