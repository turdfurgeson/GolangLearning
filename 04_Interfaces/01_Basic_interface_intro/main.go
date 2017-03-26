package main

import "fmt"

type Pilot struct {
	talk string
}

type Controller struct {
	talk string
}

type airTraffic interface {
	speak() string
}

func (p Pilot) speak() string {
	return "Request landing clearance"
}

func (c Controller) speak() string {
	return "Runway 28R cleared to land"
}

func talk(z airTraffic) {
	fmt.Println(z)
	fmt.Println(z.speak())
}

func main() {
	a := Pilot{}
	b := Controller{}
	talk(a)
	talk(b)
}

//func main() {
//	x := []airTraffic{Pilot{}, Controller{}}
//	for _, airTraffic := range x {
//		fmt.Println(airTraffic.speak())
//
//	}
//
//}
