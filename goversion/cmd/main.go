package main

import (
	"fmt"

	"github.com/stevedesilva/learngo/goversion"
	"github.com/stevedesilva/learngo/printer"
)

func main() {
	printer.Hello()
	fmt.Println(goversion.GetVersion())
}
