package main

import "fmt"

func main() {	//func loops n times and launches the go routines.
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {	//loops n times and assigns i to c
		go func() {
			for i := 0; i < n; i++ {
				c <- i
			}
			//assigns book to channel indicating done
			done <- true
		}()
	}

	go func() {	//loops and closes the channel loops above
		for i := 0; i < n; i++ {
			<-done
		}
		//closes the channel and frees it up
		close(c)
	}()

	for n := range c {	//waiting to recieve data dump from channel
		// and prints out
		fmt.Println(n)
	}
}
