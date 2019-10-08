package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {
	sumFunc()
	verboseLoop()
}

func sumFunc() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------
func verboseLoop() {
	const min, max = 1, 10
	var sum int
	for i := min; i <= max; i++ {
		sum += i
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}

	}
	fmt.Printf(" = %d\n", sum)
}
