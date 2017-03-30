package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	//same func process as the wait-group problem
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//putting value on the channel
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	//without putting the following into a go routine of their own, the
	// channels are blocked.
	go func() {
		//take value off the channel and free it up.
		<-done
		<-done
		//have to close channel for it to pull off in range below
		close(c)
	}()

	// waiting to receive block from the unbuffered channel of c.
	for n := range c {
		fmt.Println(n)
	}
}
