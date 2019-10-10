package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	// firstTurnWinner()
	// randMsg()
	doubleGuesses()
}

func firstTurnWinner() {
	const (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?
	`
	)

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxTurns; i++ {
		random := rand.Intn(maxTurns + 1)
		if guess != random {
			continue
		}
		var msg string
		if i == 1 {
			msg += "🥇 FIRST TIME"
		}
		fmt.Printf("CONGRATS YOU ARE A %s WINNER !!!\n", msg)
		return
	}
	fmt.Println("☠️  YOU LOST... Try again?")
}

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------
func randMsg() {
	const (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?
	`
	)

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxTurns; i++ {
		random := rand.Intn(maxTurns + 1)
		if guess != random {
			continue
		}
		fmt.Println(wonMsg())
		return
	}
	fmt.Println(lostMsg())
}

func lostMsg() string {
	var msg string
	switch random := getRandNum(2); random {
	case 0:
		msg = "LOSER!"
	case 1:
		msg = "YOU LOST. TRY AGAIN?"
	}
	return msg
}

func wonMsg() string {
	var msg string
	switch random := getRandNum(2); random {
	case 0:
		msg = "YOU WON"
	case 1:
		msg = "YOU'RE AWESOME"
	}
	return msg
}

func getRandNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------
func doubleGuesses() {
	const (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?
	`
	)

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxTurns; i++ {
		random := rand.Intn(maxTurns + 1)
		if guess != random && guess2 != random {
			continue
		}
		fmt.Println(wonMsg())
		return
	}
	fmt.Println(lostMsg())
}
