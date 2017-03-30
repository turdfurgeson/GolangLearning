package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//this func puts a value onto the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//closes the channel - can't put values onto channel anymore.
		// Once channel is drained/empty you can't put anything else
		// on it
		close(c)
	}()
	// this loops over the channel until it is closed or empty. It also
	// blacks the main() from exiting until there is nothing left on the
	// channel.
	for n := range c {
		fmt.Println(n)
	}
}
