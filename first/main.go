package main

import (
	"fmt"
	"os"
)

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	fmt.Println("Steve de Silva")
	fmt.Println("Clive de Silva", "Steve", "Fay")

	fmt.Println("inanc")
	fmt.Println("lina")
	fmt.Println("ebru")

}
