package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string { //fullName is the method. (p person)
	// is the receiver. The receiver attaches the function to the
	// receiver type! fullName() can be accessed by any value created by
	// type person.
	return p.first + p.last
}
func main() {
	p1 := person{"Dirk", "Diggler", 34}
	p2 := person{"Candy", "Rain", 22}
	//fmt.Println(p1.first, p1.last, p1.age)
	//fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}
