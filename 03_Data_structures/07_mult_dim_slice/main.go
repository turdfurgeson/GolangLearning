package main

import "fmt"

func main() {

	records := make([][]string, 0) //makes a multi dimensional inex of
	// strings with a length and cap of 0

	//student 1 builds student string
	student1 := make([]string, 4)
	student1[0] = "Tim"
	student1[1] = "Flores"
	student1[2] = "100.0"
	student1[3] = "90.0"

	records = append(records, student1) //adds student string slice to
	// the records slice

	//student 2 builds student string
	student2 := make([]string, 4)
	student2[0] = "Bastian"
	student2[1] = "Schweinsteiger"
	student2[2] = "65.0"
	student2[3] = "70.0"

	records = append(records, student2)

	fmt.Println(records)

}
