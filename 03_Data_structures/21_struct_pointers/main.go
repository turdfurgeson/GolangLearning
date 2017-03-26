package main

import "fmt"

type Person struct {
	first string
	last  string
	movie string
}

func main() {
	p1 := &Person{"Dirk", "Diggler", "Boogie Nights"}
	fmt.Println(p1)		//prints pointer to p1
	fmt.Printf("%T\n", p1)	//prints type
	fmt.Println(p1.first)	//can pull the value off a pointer without
	// having to put a '*' before name. Go is doing that for us ex) *p1
	// .first
	fmt.Println(p1.last)
	fmt.Println(p1.movie)

}
