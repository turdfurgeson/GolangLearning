package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	//sort package, Sort Func, sort.Reverse method, on the Interface
	// studyGroup
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
	fmt.Println("--------------------------")

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//converts s into StringSlice which is an Interface as well
	//METHODS CAN ONLY BE CALLED ON TYPES. That's why we need to convert
	// it. Then the func Sort() is called.
	sort.StringSlice(s).Sort()
	//calls reverse on the Interface interface of sort.StringSlice(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	fmt.Println("--------------------------")

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	//converts n into an IntSlice, which contains methods needed to use it
	// as Interface.
	sort.IntSlice(n).Sort()
	//calls reverse on the Interface interface sort.IntSlice(n) and also
	// returns an interface. Can use (sort.Reverse(sort.IntSlice(n))) as
	// an interface now.
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

}
