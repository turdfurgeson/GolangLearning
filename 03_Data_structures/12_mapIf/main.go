package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Jon"] = "Good Morning"
	myGreeting["Todd"] = "Buenos Dias"

	fmt.Println(myGreeting)

	//delete(myGreeting, "Todd")	//This will change the if result

	if val, exists := myGreeting["Todd"]; exists {
		fmt.Println("That value exists")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)

	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)

	}

	fmt.Println(myGreeting)
}
