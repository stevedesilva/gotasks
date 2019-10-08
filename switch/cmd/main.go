package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// timeChallenge()
	// richterScaleNum()
	// richterScaleString()
	pw()
}

func timeChallenge() {

	// fmt.Println("The hour is ", h)
	// switch true { = switch {
	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("Good evening")
	case h >= 12:
		fmt.Println("Good afternoon")
	case h >= 6:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night")
	}
}

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------
func richterScaleNum() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake")
		return
	}

	num, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	var desc string
	switch {
	case num < 2.0:
		desc = "micro"
	case num >= 2.0 && num < 3.0:
		desc = "very minor"
	case num >= 3.0 && num < 4.0:
		desc = "minor"
	case num >= 4.0 && num < 5.0:
		desc = "light"
	case num >= 5.0 && num < 6.0:
		desc = "moderate"
	case num >= 6.0 && num < 7.0:
		desc = "strong"
	case num >= 7.0 && num < 8.0:
		desc = "major"
	case num >= 8.0 && num < 10.0:
		desc = "great"
	default:
		desc = "massive"
	}
	fmt.Printf("%.2f is %s \n", num, desc)
}

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go aliens
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 8.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------
func richterScaleString() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
	}

	var desc, in string

	switch in = args[1]; in {
	case "micro":
		desc = "less than 2.0"
	case "very minor":
		desc = "2 - 2.9"
	case "minor":
		desc = "3 - 3.9"
	case "light":
		desc = "4 - 4.9"
	case "moderate":
		desc = "5 - 5.9"
	case "strong":
		desc = "6 - 6.9"
	case "major":
		desc = "7 - 7.9"
	case "great":
		desc = "8 - 9.9"
	case "massive":
		desc = "10+"
	default:
		desc = "unknown"
	}
	fmt.Printf("%s's richter scale is %s \n", in, desc)
}

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func pw() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	// switch {
	// case u != user && u != user2:
	// 	fmt.Printf(errUser, u)
	// case u == user && p == pass:
	// 	// fmt.Printf(accessOK, u)
	// 	fallthrough
	// case u == user2 && p == pass2:
	// 	fmt.Printf(accessOK, u)
	// default:
	// 	fmt.Printf(errPwd, u)
	// }

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		// notice this one (no more duplication)
		fallthrough
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}

	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }
}
