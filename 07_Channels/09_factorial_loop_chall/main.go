package main

import (
	"fmt"
)

func main() {
	in := gen()
	f := factorial(in)
	// n ranges over the channel in f which therefore is pulling the
	// values off of the func factorial return
	for n := range f {
		fmt.Println(n)
	}
}
//returns a channel which func factorial receives
func gen() <-chan int {
	//opens channel to receive
	out := make(chan int)
	go func() {
		//loop for 10 times
		for i := 0; i < 10; i++ {
			//computes factorial of 3 through 12 which  factors
			// 10 times. 10 x 10 is 100 computations.
			for j := 3; j < 13; j++ {
				//puts 100 numbers onto channel
				out <- j
			}
		}
		close(out)
	}()
	return out

}
//receives a channel from func gen, and then returns a channel. Essentially
// it is pulling off the value of in!
func factorial(in <-chan int) <-chan int {
	// opens channel to receive
	out := make(chan int)
	go func() {
		//value being pulled off channel through range.
		for n := range in {
			//for each of the 100 numbers received from in, it
			// calls fact and puts value onto channel
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	sum := 1
	for i := 1; i <= n; i ++ {
		sum *= i
	}
	return sum
}