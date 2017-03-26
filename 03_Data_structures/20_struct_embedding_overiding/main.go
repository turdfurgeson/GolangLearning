package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) talk() {
	fmt.Println("Hello, darlin! My name is " + p.Name)

}

type Android struct {
	Person //Person is "embedded in Android and therefore
	// carries with it the fields from the Person struct.
	Model string
	Name string	//You can have same field as the inner type. But the
	// out-type always overrides the innertype unless specified(p.Person
	// .Name) It is the same if the outer type has a method the same name
	// as the inner-type.
}

func main() {
	a := Android{
		Person: Person{ //field: type {define} - good
			// practice to do it this way with embedding.
			Name: "Charlotte",
			Age:  27,
		},
		Model: "Hitachi",
	}
	fmt.Println("-----------------------------------------------------")
	a.talk()
	fmt.Println("I am " + fmt.Sprint(a.Age) + " years old.")
	//fmt.Sprint.(i) is a slower but easier way to convert int to string.
	fmt.Println("I was built by the " + a.Model + " Corporation.")
	fmt.Println("-----------------------------------------------------")
}
