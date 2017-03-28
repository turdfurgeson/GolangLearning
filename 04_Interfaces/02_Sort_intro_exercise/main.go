package main

import (
	"fmt"
	"sort"
)

type people []string

//These methods are attached to type people.Should trigger sort package used
// here since you need these three methods to use the Interface interface!
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	//studyGroup is actually Interface interface since is uses the above
	// 3 methods
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	//Uses the Strings method from the sort package. No interface needed.
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)
	//Another way to do it. StringSlice includes the above methods which
	// means it can me implemented in the Interface interface type.
	// Basically means that anywhere Interface is implemented in the sort
	// package, you can use sort.StringSlice(s) to pass in.
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	//This way implements the type Interface being passed into the Sort
	// func. https://godoc.org/sort#Sort
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}
