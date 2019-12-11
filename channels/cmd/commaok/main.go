package main

import "fmt"

func main() {
	commaOkIdiom()
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
