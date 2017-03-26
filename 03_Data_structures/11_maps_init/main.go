package main

import "fmt"

func main() {

	map1 := make(map[string]int) //defines map1 and tells it the key
	// are strings and the value are ints

	map1["k1"] = 1 //sets key to string k1 and value to int 1
	map1["k2"] = 2

	fmt.Println("map:", map1)
	fmt.Println("len: ", len(map1)) //prints length of map

	v1 := map1["k1"] //assigns the VALUE of key k1 of map1 to var v1
	fmt.Println("v1: ", v1)

	delete(map1, "k2") //deletes key k2 and its value!
	fmt.Println("map:", map1)
	fmt.Println("len: ", len(map1))

	v, ok := map1["k2"] //checks to see if key is still in map or not
	// . returns a bool. "comma/ok idiom. Tests to see if value is there.
	// If it is, you can get it out. v is set to return the value in the
	// print statement
	fmt.Println("Is present: ", ok, v)

	map2 := map[string]int{"foo": 3, "bar": 4} //composite literal
	// form of making a map
	fmt.Println("map2:", map2)
	fmt.Println("key foo:", map2["foo"])

	var myGreeting = make(map[string]string)	//Just a different
	// way to init a map

	// var myGreeting map[string]string{
	// 	"Tim":		"Hello!"
	//	"Jenny":	"Good day!"
	// }

	// This way is just a different way to initialize empty map without a
	// nil value. Without the {} it will only value nil. With a slice you
	// can get out of that with the 'append' func. In map you can't. {}
	// can also be empty and you just update it like below.
	myGreeting["Jon"] = "Good Morning"	//this just updates map with
	// new key/value pair
	myGreeting["Todd"] = "Buenos Dias"
	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))	//prints length

}
