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
	firstTurnWinner()
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
