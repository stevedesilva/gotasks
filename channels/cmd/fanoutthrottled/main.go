package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fanOutExample()
}

func fanOutExample() {
	// general channels
	c1 := make(chan int)
	c2 := make(chan int)

	// start goroutine
	go populate(c1)

	go fanOutIn(c1, c2)

	// range over c2 channel until closed - blocks waiting for value
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("The end!!!!")
}

// specific send channels
func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines) // limit goroutines

	// fanout - launch 10 goroutines which will loop over c
	// then fanin (write back) to the same c2 channel
	// range exits after c1 closes in populate
	for i := 0; i < goroutines; i++ {
		// each goroutine
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()

		}()
	}
	// wait for all goroutines to complete
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
