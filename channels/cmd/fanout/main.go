package main

import (
	"fmt"
	"sync"
)

func main() {
	sendAndReceive()
}

func sendAndReceive() {
	// general channels
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// start goroutines
	go send(eve, odd)

	go receive(eve, odd, fanin)

	// range over fanin channel until closed - blocks waiting for value
	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("The end!!!!")
}

// specific send channels
func send(eve, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	// close channel stop range
	close(eve)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)
	// put multiple values from even and odd into one chan (fanin). Reduce to 1 channel
	go func() {
		// range until the channel is closed by send go routine
		// block until value is pull off chan
		for v := range even {
			fanin <- v
		}
		// decrement waitgroup count
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		// decrement waitgroup count
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
