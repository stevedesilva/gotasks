package main

import (
	"fmt"

	"github.com/stevedesilva/learngo/printer"
)

func main() {
	printer.Hello()
	p1 := printer.NewPerson("Steve", 41)
	fmt.Println(p1)
	fmt.Println(p1.GetPersonName())
	fmt.Println(p1.GetPersonAge())

	fmt.Println(-0, 1., .2, true, 1, 2)
	hexPrinter()

}

func hexPrinter() {

	for i := 0x0; i <= 0x9; i++ {
		fmt.Println(i)
	}

	for i := 0xa; i <= 0xf; i++ {
		fmt.Println(i)
	}

}
