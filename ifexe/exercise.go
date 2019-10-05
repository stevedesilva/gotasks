package ifexe

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Q1 exercise
// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------
//
func Q1() {
	// Change this accordingly to produce the expected outputs
	age := 10

	// Type your if statement here.
	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Adulthood")
	} else {
		fmt.Println("Booting up")
	}
}

// Q2 exe
// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------
func Q2() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if radius >= 200 && isSphere {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

// Q3 exe
// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go i wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------
func Q3() {

	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q.\n", args[1])
	} else if l == 2 {
		fmt.Printf("There are two: %q %q.\n", args[1], args[2])
	} else {
		fmt.Printf("There are %d arguments.\n", l)
	}
}

// F4 exe
// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------
func Q4() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 || l > 1 || len(args[1]) > 1 {
		fmt.Println("Give me a letter")
		return
	}

	arg := args[1]
	if i := strings.IndexAny(args[1], "aeiou"); i != -1 {
		fmt.Printf("%q is a vowel. \n", arg)
	} else if i := strings.IndexAny(args[1], "wy"); i != -1 {
		fmt.Printf("%q is sometimes a vowel, sometimes not. \n", arg)
	} else if i := strings.IndexAny(args[1], "bcdfghjklmnpqrstvxz"); i != -1 {
		fmt.Printf("%q is a consonant. \n", arg)
	} else {
		fmt.Println("Give me a letter")
	}

	// //strings.Size
	// if l == 0 {
	// 	fmt.Println("Give me a letter")
	// } else if al := len(args[1]); l == 1 && al > 1 {
	// 	fmt.Println("Give me a letter not a word ")
	// } else if al := len(args[1]); l == 1 && al == 1 {
	// 	i := strings.IndexAny(args[1], "aeiou")
	// 	fmt.Println(i)

	// 	if i != -1 {
	// 		fmt.Printf("%s is a vowel.", al)

	// 	}
	// 	j := strings.IndexAny(args[1], "wy")
	// 	if j != -1 {
	// 		fmt.Printf("%q is sometimes a vowel, sometimes not.", j)
	// 	}
	// 	k := strings.IndexAny(args[1], "bcdfghjklmnpqrstvxz")
	// 	if j != -1 {
	// 		fmt.Printf("%q is a consonant.", k)

	// 	}

	// } else if l > 2 {
	// 	fmt.Println("Give me a letter not more than 1 word ")
	// } else {
	// 	fmt.Println("Other ")
	// }

}

// Q5 exe
// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------
func Q5() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Requires age")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: Something went wrong %v .\n", err)
	}

	if n < 0 {
		fmt.Printf("Wrong age: %q.\n", args[1])
	} else if n <= 13 {
		fmt.Println("PG-Rated")
	} else if n <= 17 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("R-Rated")
	}
}

// Q6 exe
// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------
func Q6() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Pick a number")
		return
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a number.\n", args[1])
		return
	}

	if num%2 == 0 {
		fmt.Printf("%d is an even number", num)
		if num%8 == 0 {
			fmt.Printf(" and it's divisible by 8")
		}
		fmt.Println("")

	} else {
		fmt.Printf("%d is an odd number.\n", num)

	}

}

// FQ exe
// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------
func Q7() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println(" Give me a year number")
		return
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", args[1])
		return
	}

	if n%4 == 0 && (n%100 != 0 || n%400 == 0) {
		fmt.Printf("%d is a leap year.\n", n)
	} else {
		fmt.Printf("%d is not a leap year.\n", n)
	}

}

// Q8 exe
// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------
func Q8() {
	// type mdays struct {
	// 	month string
	// 	days int
	// }

	// months := []mdays {
	// 	{"january",31},
	// 	{"february",28},
	// 	{"march",31},
	// 	{"april",31},
	// 	{"may",31},
	// 	{"june",30},
	// 	{"july",31},
	// 	{"august",31},
	// 	{"september",30},
	// 	{"october",31},
	// 	{"november",30},
	// 	{"december",31},
	// }

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	mname := args[1]
	switch m := strings.ToLower(args[1]); m {
	case "january", "march", "april", "may", "july", "august", "october", "december":
		fmt.Printf("%q has 31 days.\n", mname)
	case "february":
		var lyr bool
		if n := time.Now().Year(); n%4 == 0 && (n%100 != 0 || n%400 == 0) {
			fmt.Printf("%d is a leap year.\n", mname)
			lyr = true
		} else {
			fmt.Printf("%d is not a leap year.\n", mname)
			lyr = false
		}
		if lyr {
			fmt.Printf("%q has 29 days.\n", mname)
		} else {
			fmt.Printf("%q has 28 days.\n", mname)
		}

	case "june", "september", "november":
		fmt.Printf("%q has 30 days.\n", mname)
	default:
		fmt.Printf("%q is not a month.\n", mname)
	}
}
