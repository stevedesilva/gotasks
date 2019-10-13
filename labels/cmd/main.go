package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------
const (
	corpus = "lazy cat jumps again and again and again"
)

func main() {
	caseInSearch()
}

func caseInSearch() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter a word.")
		return
	}
	w := strings.ToLower(args[0])

	fields := strings.Fields(corpus)

	for i, v := range fields {
		if w == v {
			fmt.Printf("#%-2d: %q\n", i+1, w)
		}
	}

}
