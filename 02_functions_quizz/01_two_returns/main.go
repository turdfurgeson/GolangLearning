package main

import "fmt"

func half(x int) (int, bool) {

	return x / 2, x%2 == 0

}

func main() {

	a, even := half(8)

	fmt.Println(a, even)

}
