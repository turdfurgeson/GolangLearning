package main

import (
	"fmt"
	"time"
)

func main() {
	//unbuffered channel that takes an int.
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {	//channel is passed to the Print line.
			fmt.Println(<-c)

		}
	}()
	time.Sleep(time.Second)
}

//Think of channels like relay races. The baton is being held by the runner
// until it is passed on to the next runner. Same for channels-the value is
// being held in the channel until being passed to something.
