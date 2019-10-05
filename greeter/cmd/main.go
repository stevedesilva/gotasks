package main

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------
import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// UNCOMMENT & FIX THIS CODE
	// lengthArgs()
	// printPath()
	// printYourName()
	greetPeople()
}

func lengthArgs() {
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
}

/// STEPS:
//
// Compile it by typing:
//   go build -o myprogram
//
// Then run it by typing:
//   ./myprogram
//
// If you're on Windows, then type:
//   myprogram
// -------------------------------------------------------
func printPath() {
	fmt.Println(os.Args[0])
}

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func printYourName() {
	// get your name from the command-line
	// and print it
	name := os.Args[1]
	fmt.Printf("Hi %v\n How are you?", name)
}

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match to the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store the all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func greetPeople() error {
	// TYPE YOUR CODE HERE

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
	length := len(os.Args) - 1
	if length < 3 {
		return errors.New("Not enough args")
	}
	fmt.Printf("There are %d people! \n", length)
	for i, j := range os.Args {
		if i > 0 {
			fmt.Printf("Hello great %v! \n", j)
		}

	}
	fmt.Println("Nice to meet you all.")
	return nil
}
