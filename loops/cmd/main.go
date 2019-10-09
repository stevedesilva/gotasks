package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

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
	// sumFunc()
	// verboseLoopFixed()
	// verboseLoopCmdLine()
	// onlyEven()
	// infiniteSum()
	infiniteWait()
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
func verboseLoopFixed() {
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

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func verboseLoopCmdLine() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide min and max values.")
		return
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please provide min value.")
		return
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Please provide max value.")
		return
	}
	if max < min {
		fmt.Println("Max cannot be less than min")
		return
	}

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

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------
func onlyEven() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide min and max values.")
		return
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please provide min value.")
		return
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Please provide max value.")
		return
	}
	if max < min {
		fmt.Println("Max cannot be less than min")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}

	}
	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func infiniteSum() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide min and max values.")
		return
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please provide min value.")
		return
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Please provide max value.")
		return
	}
	if max < min {
		fmt.Println("Max cannot be less than min")
		return
	}

	var (
		i   = min
		sum int
	)
	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Infinite Kill
//
//  1. Create an infinite loop
//  2. On each step of the loop print a random character.
//  3. And, sleep for 250 milliseconds.
//  4. Run the program and wait a couple of seconds
//     then kill it using CTRL+C keys
//
// RESTRICTIONS
//  1. Print one of those characters randomly: \ / | -
//  2. Before printing a character print also this
//     escape sequence: \r
//
//     Like this: "\r/", or this: "\r|", and so on...
//
//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
//
// HINTS
//  1. Use `time.Sleep` to sleep.
//  2. You should pass a `time.Duration` value to it.
//  3. Check out the Go online documentation for the
//     `Millisecond` constant.
//  4. When printing, do not use a newline! Or a Println!
//     Use Print or Printf instead.
//
// NOTE
//  If this exercise is hard for you now, wait until the
//  lucky number lecture. Even then so, then just skip it.
//
//  You can return back to it afterwards.
//
// EXPECTED OUTPUT
//  - The program should display the following messages in any order.
//  - And, the first character (\, -, /, or |) should be randomly
//    displayed.
//  - \r or \f sequence will clear the line.
//
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ---------------------------------------------------------

func infiniteWait() {

	for {
		var c string
		switch rand.Intn(4) {

		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}

		fmt.Printf("\r%s Please Wait. Processing....\n",c)
		time.Sleep(time.Millisecond * 250)
	}
}
