package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/stevedesilva/learngo/gotype_sys/example"
)

// ---------------------------------------------------------
// EXERCISE: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------

func q1() {
	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var letter byte = 'a'
	fmt.Println(" an english letter", letter)
	// a non-english letter (search web for: unicode codepoint)
	var unicode rune
	unicode = 'é'
	fmt.Println(" a non english letter", unicode)
	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println(" a  year", year)
	// a month in 2 digits: 1 to 12
	var month uint8 = 1
	fmt.Println(" a  month", month)

	// the speed of the light
	var lightSpeed int32 = 670616629 // miles
	fmt.Println("the speed of the light:", lightSpeed)

	// // angle of a circle
	// var circle uint16 = 360
	// fmt.Println(" a  circle", circle)
	// // PI number: 3.141592653589793
	// var pi float32 = 3.141592653589793
	// fmt.Println(" pi", pi)

	// angle of a circle
	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)

	var pi float64
	pi = 3.141592653589793
	fmt.Println("PI number:", pi)
}

// ---------------------------------------------------------
// EXERCISE: The Type Problem
//
//  Solve the data type problem in the program.
//
// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true
// ---------------------------------------------------------

func q2() {
	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.
	var (
		width  uint16
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// UNCOMMENT THIS:
	fmt.Println("are they equal?", width == height)
}

// ---------------------------------------------------------
// EXERCISE: Parse Arg Numbers
//
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.
//
// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------

func q3() {
	// --------------------------------------
	// EXAMPLE:
	// --------------------------------------
	// How to get an int8 from command-line:
	// First argument should be a value of -128 to 127 range
	//
	// Second argument: 10 means decimal number
	// Third argument : 8 means 8-bits (int8)
	input := os.Args[1]
	val, _ := strconv.ParseInt(input, 10, 8)

	// Now the val variable is int64 because ParseInt
	// returns an int64. But, since I passed 8 as its third
	// argument, it returns int8 convertable value.
	//
	// Try running the program with a value of -128 to 127
	// Running it beyond that range will result in
	// either -128 or 127.
	fmt.Println("int8 value is:", int8(val))

	// --------------------------------------
	// NOW IT'S YOUR TURN!
	// --------------------------------------

	// 1. Get an int16 value using ParseInt and convert it and print it
	val16, _ := strconv.ParseInt(input, 10, 16)
	fmt.Println("int16 value is:", int16(val16))
	// 2. Get an int32 value using ParseInt and convert it and print it
	val32, _ := strconv.ParseInt(input, 10, 32)
	fmt.Println("int32 value is:", int32(val32))
	// 3. Get an int64 value using ParseInt and convert it and print it
	val64, _ := strconv.ParseInt(input, 10, 64)
	fmt.Println("int64 value is:", int64(val64))
	// 4. Get an int8 value using ParseInt and convert it and print it
	//    But this time, get the value in bits.
	//
	//    For example : 00000100
	//    Should print: 4
	fmt.Printf("%8b", val)

	// 5th argument is a number in bits. And its int8.
	// ParseInt(.., 2, ...) -> 2 means binary base
	val5, _ := strconv.ParseInt(os.Args[2], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[2], int8(val5))
}

// ---------------------------------------------------------
// EXERCISE: Time Multiplier
//
//  You should get an argument from the command-line and
//  you need to multiply the time duration value `t` with
//  the given argument.
//
//  1. Get an argument from the command-line
//  2. Convert it to int64 and store it in a variable
//  3. Multiply the `t` variable with that variable
//  4. Print the result
//
// HINT
//  You can use ParseInt to convert the command-line
//    argument into an int64 value.
//
//  You can skip the error value using a blank-identifier.
//
// EXPECTED OUTPUT
//
//  When runned like this:
//    go run main.go 2
//
//  It should print this:
//    3h0m0s
// ---------------------------------------------------------

func q4() {
	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	// TYPE YOUR CODE HERE
	in := os.Args[1]
	multiplier, _ := strconv.ParseInt(in, 10, 64)

	t *= time.Duration(multiplier)
	fmt.Printf("%d\n", multiplier)

	// DONT TOUCH THIS
	fmt.Println(t)
}

// ---------------------------------------------------------
// EXERCISE: Refactor Feet to Meter
//
//  Remember the feet to meters program?
//  Now, it's time to refactor it.
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------

func q5() {
	// ----------------------------
	// 1. Define Feet and Meters types below
	//    Their underlying type can be float64
	// type Feet float64
	// type Meters float64
	type (
		Feet   float64
		Meters float64
	)

	// ----------------------------
	// 2. UNCOMMENT THE CODE BELOW THEN DON'T TOUCH IT
	var (
		feet   Feet
		meters Meters
	)

	// ----------------------------
	// 3. Get feet value from the command-line

	// 4. Convert it to an float64 first using ParseFloat
	val, _ := strconv.ParseInt(os.Args[1], 10, 64)

	// 5. Then, convert it into a Feet type
	feet = Feet(val)

	// ----------------------------
	// 6. Uncomment the code below
	// 7. And, convert the expression to Meters type

	meters = Meters(feet * 0.3048)

	// ----------------------------
	// 8. UNCOMMENT THE CODE BELOW
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}

// ---------------------------------------------------------
// EXERCISE: Convert the Types
//
//  Convert the variables to appropriate types.
//
// EXPECTED OUTPUT
//  325.5 299.5
// ---------------------------------------------------------

func q6() {
	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	// ----------------------------------------
	// FIX THE CODE BELOW:
	// You should solve it only by using conversions.
	// Do not change the code in any other way.

	// celsius *= Celsius(celsiusDegree * Temperature(factor))
	// fahr *= Fahrenheit(fahrDegree * Temperature(factor))
	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	// ----------------------------------------
	// DONT TOUCH THIS
	fmt.Println(celsius, fahr)

	// // YOU MAY REMOVE THESE WHEN YOU'RE DONE
	// _, _, _ = celsiusDegree, fahrDegree, factor
}

func main() {
	// q1()
	// q2()
	// q3()
	// q4()
	// q5()
	// q6()
	example.Conversion()

}
