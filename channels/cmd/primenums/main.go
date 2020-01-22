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
	fmt.Printf("Var ch :  %v \n",ch)
	defer close(ch)

	// keeps its number original chan
	go generate(ch)

	for i:=0; i<maxNumber;i++{
		prime := <-ch
		fmt.Println("prime=", prime," i=",i)
		ch1:=make(chan int)
		go filter(ch,ch1,prime)
		ch=ch1
	}
}

func generate(ch chan<- int){
	fmt.Printf("generate start ch :  %v \n",ch)
	for i:=2; ;i++ {
		ch <-i
	}
}

func filter(in <-chan int,out chan <-int, prime int){
	fmt.Printf("filter in(%v) out(%v) p(%d()) \n",in,out,prime)
	for{
		i:= <-in
		if i%prime !=0{
			out <- i
		}
	}
}