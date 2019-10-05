package exercise

import (
	"fmt"
	"time"
)

// Q1  exercise
// ---------------------------------------------------------
// EXERCISE: Minutes in Weeks
//
//  Calculate how many minutes in two weeks.
//
//  STEPS:
//  1. Declare `minsPerDay` constant and initialize it
//     to the number of minutes in a day
//
//  2. Declare `weekDays` constant and initialize it
//     to 7.
//
//  3. Use printf to print the total number of minutes
//     in two weeks.
//
// EXPECTED OUTPUT
//  There are 20160 minutes in 2 weeks.
// ---------------------------------------------------------
func Q1() {
	const (
		minsPerDay, weekDays = 1440, 7
	)
	twoWeeks := 2
	minsIn2weeks := minsPerDay * weekDays * twoWeeks

	fmt.Printf("There are %d minutes in %d weeks. \n", minsIn2weeks, twoWeeks)
}

// Q2 exercise
// ---------------------------------------------------------
// EXERCISE: Remove the Magic
//
//  Get rid of the magic numbers in the following code.
//
// RESTRICTIONS
//  1. You should declare 3 constants named:
//       hoursInDay, daysInWeek, and hoursPerWeek
//
//  2. And, hoursPerWeek constant should be initialized
//     using hoursInDay and daysInWeek constants.
//
// EXPECTED OUTPUT
//  There are 840 hours in 5 weeks.
// ---------------------------------------------------------
// Q2
func Q2() {

	// const (
	// 	hoursInDay   = 24
	// 	daysInWeek   = 7
	// 	hoursPerWeek = hoursInDay * daysInWeek
	// )

	// fmt.Printf("There are %d hours in %d weeks.\n",
	// 	hoursPerWeek*5, 5)

	const (
		hoursInDay   = 24
		daysInWeek   = 7
		fiveWeeks    = 5
		hoursPerWeek = hoursInDay * daysInWeek
	)
	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*fiveWeeks, fiveWeeks)
}

// Q3 exe
// ---------------------------------------------------------
// EXERCISE: Constant Length
//
//  Calculate how many characters inside the `home`
//  constant and print it.
//
// STEPS:
//  1. Declare a constant named `home`
//  2. Initialize it to "Milky Way Galaxy" string literal
//
//  3. Declare another constant named `length`
//  4. Initialize it by using the built-in function `len`.
//
//  5. Print the message below using the constants that
//     you've declared.
//
// RESTRICTION:
//  Use Printf.
//  Print the `home` constant using the quoted-string verb.
//
// EXPECTED OUTPUT
//  There are 16 characters inside "Milky Way Galaxy"
// ---------------------------------------------------------
func Q3() {
	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q.\n", length, home)

}

// Q4 exe
// ---------------------------------------------------------
// EXERCISE: TAU
//
//  Fix the following program and print the TAU number.
//
// HINT
//  You can use %g verb for printing tau.
//
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------
func Q4() {
	// What's the problem with this code?
	// Why it doesn't work?

	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)
	// const pi, tau = 3.14159265358979323846264, pi * 2
	fmt.Printf("tau = %g \n", tau)

}

// Q5 exe
// ---------------------------------------------------------
// EXERCISE: Area
//
//  Fix the following program.
//
// RESTRICTION
//  You should not use any variables.
//
// EXPECTED OUTPUT
//  area = 1250
// ---------------------------------------------------------
func Q5() {
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}

// Q6 exe
// ---------------------------------------------------------
// EXERCISE: No Conversions Allowed
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------
func Q6() {
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}

// Q7 exe
// ---------------------------------------------------------
// EXERCISE: Iota Months
//
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
//
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
//
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------
func Q7() {
	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)
}

// Q8 exe
// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------
func Q8() {
	// HINT: This is a valid constant declaration
	//       Blank-Identifier can be used in place of a name
	const _ = iota
	//    ^- this is just a name

	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}

// Q9 exe
// ---------------------------------------------------------
// EXERCISE: Iota Seasons
//
//  Use iota to initialize the season constants.
//
// HINT
//  You can change the order of the constants.
//
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------
func Q9() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		// _ =  iota * 3
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}


