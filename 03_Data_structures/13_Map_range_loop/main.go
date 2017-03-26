package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
		2: "You're awesome and hello",
		3: "Love your pants. What's your name?",
		4: "You ain't right. But what's up anyways?",
	}

	for key, value := range myGreeting { //any interable object you
		// can use a range. key/val are just variable names. they cna mean
		// anything you want.
		fmt.Println(key, " - ", value)
	}
}
