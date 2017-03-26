package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Dirk", "Diggler", 34}
	p2 := person{"Candy", "Rain", 22}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
