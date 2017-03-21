package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {

		return x / 2, x%2 == 0

	}

	a, even := half(8)

	fmt.Println(a, even)

}
