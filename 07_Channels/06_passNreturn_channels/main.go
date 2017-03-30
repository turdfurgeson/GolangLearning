package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out

}

func puller(c chan int) chan int {
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		//have to close channel if you use range.
		close(out2)
	}()
	return out2
}
