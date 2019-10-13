package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
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
	// doubleGuesses()
	// verboseMode()
	enoughPicks()
}

func firstTurnWinner() {

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
		random := rand.Intn(guess + 1)
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
		random := rand.Intn(guess + 1)
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

	min := guess
	if guess < guess2 {
		min = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(min + 1)

		if n == guess || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------
func verboseMode() {
	const (
		usage    = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?
	(Provide -v flag to see the picked numbers.)
	`
	)
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool
	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.", guess)
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxTurns; i++ {
		random := rand.Intn(guess + 1)
		if verbose {
			fmt.Printf("%d ", random)
		}

		if guess != random {
			continue
		}
		fmt.Println(wonMsg())
		return
	}
	fmt.Println(lostMsg())
}

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func enoughPicks() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	random := 10
	if guess > random {
		random = guess
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= maxTurns; i++ {

		if guess == rand.Intn(random+1) {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}
	fmt.Println("☠️  YOU LOST... Try again?")

}

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------
func dynamicDifficulty() {

}
