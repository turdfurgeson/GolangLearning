package main

import (
	"fmt"
)


func foo(num ...int) int {

	var total int

	for _, y := range num {

		total += y
	}

	return total
}

func main () {

	fmt.Println(foo(1,2))
	fmt.Println(foo(1,2, 3, 4))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())


}