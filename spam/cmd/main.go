package main

import (
	"fmt"
	"os"

	"github.com/stevedesilva/learngo/spam"
)

func main() {

	s := spam.New()
	fmt.Printf("Type %T value %[1]v \n", s)

	res, err := s.Process(os.Args)
	fmt.Printf("Result\t%v\nError\t%v\n", res, err)

}
