package main

import (
	"fmt"
	"sync"
)

func main() {
	//step one: declare channel and wait group
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	//step 2: launch the next three functions.
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	//step 3: rec ieve from channel c.
	for n := range c {
		fmt.Println(n)
	}
}

/* A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can
 be used to block until all goroutines have finished. https://golang
 .org/pkg/sync/#WaitGroup */