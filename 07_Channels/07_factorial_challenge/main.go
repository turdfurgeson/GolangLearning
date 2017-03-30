package main

import (
	"fmt"
)

func main() {
	f := factorial(5)
	for x := range f {
		fmt.Println(x)
	}
}

func factorial(x int) chan int {
	out := make(chan int)

	go func() {
		sum := 1
		for i := 1; i <= x; i++ {
			sum *= i
		}
		out <- sum
		close(out)

	}()
	return out
}
