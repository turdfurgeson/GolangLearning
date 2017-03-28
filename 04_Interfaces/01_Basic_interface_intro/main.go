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
	return p.talk
}

func (c Controller) speak() string {
	return c.talk
}

func talk(z airTraffic) {
	//fmt.Println(z)
	fmt.Println(z.speak())
}

func main() {
	a := Pilot{"Request landing clearance"}
	b := Controller{"Runway 28R cleared to land"}
	talk(a)
	talk(b)
}
