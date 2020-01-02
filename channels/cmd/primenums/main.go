package main

import (
	"runtime"
	"fmt"
)

func main() {
	nCPU:=runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)


	const maxNumber  = 100

	ch := make(chan int)
	defer close(ch)

	go Generate(ch)

	for i:=0; i<maxNumber;i++{
		prime := <-ch
		fmt.Println(prime)
		ch1:=make(chan int)
		go Filter(ch,ch1,prime)
		ch=ch1
	}
}

func Generate(ch chan<- int){
	for i:=2; ;i++ {
		ch <-i
	}
}

func Filter(in <-chan int,out chan <-int, prime int){
	for{
		i:= <-in
		if i%prime !=0{
			out <- i
		}
	}
}