package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	const maxNumber = 3
	// in channel
	ch := make(chan int)

	fmt.Printf("Var ch :  %v \n", ch)
	defer close(ch)

	// keeps hold of ref number original chan
	go generate(ch)

	fmt.Printf("There were %v active goroutines, Init \n", runtime.NumGoroutine())

	for i := 0; i < maxNumber; i++ { // loop 100 times
		prime := <-ch // read from in 2, gen() in <- 3
		fmt.Printf("main in(%v) prime(%d) i(%d) \n", ch, prime, i)
		out := make(chan int) // create new out stores
		go filter(ch, out, prime)
		ch = out // replace in ch with out
	}
	// // Wait for the next tick to ensure the goroutines have been cleaned up
	time.Sleep(2 * time.Millisecond)
	fmt.Printf("There were %v active goroutines, End \n", runtime.NumGoroutine())

}

// new gr
func filter(in <-chan int, out chan<- int, prime int) {
	fmt.Printf("filter init: in(%v) out(%v) p(%d)\n", in, out, prime)
	for {
		i := <-in // first time get origin in ch 3, gen() in <-4
		fmt.Printf("filter     : in(%v) out(%v) i(%d)\n", in, out, i)
		if i%prime != 0 {
			fmt.Printf("filter >>> : in(%v) out(%v) i(%d)\n", in, out, i)
			out <- i
		}

	}
}

func generate(ch chan<- int) {
	// fmt.Printf("generate start ch :  %v \n", ch)
	for i := 2; ; i++ {
		// fmt.Printf("generate ch :  %d \n", i)

		ch <- i
	}

}

// Var ch :  0xc0000b6000 
// There were 2 active goroutines, Init 

// main in(0xc0000b6000) prime(2) i(0) 
// filter init: in(0xc0000b6000) out(0xc0000be000) p(2)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(3)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(3)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(4)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(5)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(5)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(6)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(7)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(7)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(8)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(9)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(9)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(10)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(11)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(11)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(12)
// filter     : in(0xc0000b6000) out(0xc0000be000) i(13)
// filter >>> : in(0xc0000b6000) out(0xc0000be000) i(13)

// main in(0xc0000be000) prime(3) i(1) 
// filter init: in(0xc0000be000) out(0xc0000be060) p(3)
// filter     : in(0xc0000be000) out(0xc0000be060) i(5)
// filter >>> : in(0xc0000be000) out(0xc0000be060) i(5)
// filter     : in(0xc0000be000) out(0xc0000be060) i(7)
// filter >>> : in(0xc0000be000) out(0xc0000be060) i(7)
// filter     : in(0xc0000be000) out(0xc0000be060) i(9)
// filter     : in(0xc0000be000) out(0xc0000be060) i(11)
// filter >>> : in(0xc0000be000) out(0xc0000be060) i(11)

// main in(0xc0000be060) prime(5) i(2) 
// filter init: in(0xc0000be060) out(0xc000080060) p(5)
// filter     : in(0xc0000be060) out(0xc000080060) i(7)
// filter >>> : in(0xc0000be060) out(0xc000080060) i(7)




// There were 5 active goroutines, End 