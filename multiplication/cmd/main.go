package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const max = 5

func main() {
	// multiexe()
	// e1()
	// e1Multiplication()
	// e2Algebra()
	fmt.Println(2 % 3)
}

func multiexe() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give number for the table")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

func e1Multiplication() {
	args := os.Args[1:]
	if l := len(args); l != 1 {
		fmt.Println("Give me the size of the table")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil || num < 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= num; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for j := 0; j <= num; j++ {
		fmt.Printf("%5d", j)
		for k := 0; k <= num; k++ {
			fmt.Printf("%5d", j*k)
		}
		fmt.Println()
	}

}

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------
func e2Algebra() {

	const (
		ops      = "%*/+-"
		usage    = "Usage: op=" + ops + "[size]"
		validOps = `Invalid operator.
		Valid ops one of: ` + ops
		sizeMissingMsg = "Size is missing"
		wrongSizeMsg   = "Wrong size"
	)

	args := os.Args[1:]
	if l := len(args); l != 2 {
		if l == 1 {
			fmt.Println(sizeMissingMsg)
		}
		fmt.Println(usage)
		return
	}

	sym := args[0]
	if i := strings.IndexAny(sym, ops); i == -1 {
		fmt.Println(validOps)
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil || num < 0 {
		fmt.Println(wrongSizeMsg)
		return
	}

	fmt.Printf("%5s", sym)
	for i := 0; i <= num; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for j := 0; j <= num; j++ {
		fmt.Printf("%5d", j)
		for k := 0; k <= num; k++ {
			fmt.Printf("%5d", calcResult(j, k, sym))
		}
		fmt.Println()
	}
}

func calcResult(j, k int, sym string) int {
	var res int

	switch sym {
	case "*":
		res = j * k

	case "-":
		res = j - k

	case "+":
		res = j + k

	case "/":
		if k != 0 {
			res = j / k
		}

	case "%":
		if k != 0 {
			res = j % k
		}

	}
	return res
}
