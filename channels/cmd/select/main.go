package main

import "fmt"

func main() {
	sendAndReceive()
	commaOkIdiom()
}

func sendAndReceive() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("The end!!!!")
}

func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	// close(eve)
	// close(odd)
	close(quit)
}

func receive(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("from the eve channel:", v)
		case v := <-odd:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from the quit channel:", i, ok)
				return
			}
			fmt.Println("from the quit channel:", i, ok)
		}
	}
}

func commaOkIdiom() {
	fmt.Println("Hello, playground")

	in := make(chan int)
	go func(in chan<- int) {
		in <- 42
		close(in)
	}(in)

	out, ok := <-in
	fmt.Println("Out: ", out, ok)
	out, ok = <-in
	fmt.Println("Out 2: ", out, ok)

}
