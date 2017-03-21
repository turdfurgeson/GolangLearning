package main

import "fmt"

func great(num ...int) int {

	var high int

	for _, v := range num {

		if v > high {

			high = v

		}

	}

	return high

}

func main() {
	max := great(4, 6, 7, 8, 123, 5, 34, 765)
	fmt.Println(max)
}
