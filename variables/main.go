package main

import "fmt"

var nocomplilerror string // package scope variable may be used else where - hence no compile error if not used

func main() {
	speedFunc()
	printDefaults()
	shortDec()
	multShortDec()
	sumInit()
	shortDiscard()
	redeclareFunc()
	assignmentFunc()
	assignmentFunc2()
	multiCaller()
}

func speedFunc() {
	var speed int
	var heat float64
	var off bool
	var brand string
	var compilerror string
	_ = compilerror // blank identifier (stops compile error)
	var height int

	fmt.Println(speed, height)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)

	// var 3speed int
	// var !speed int
	// var spe?ed int
	// var var int
	// var func int
	// var package int
}

func printDefaults() {
	//    int
	//    int8
	//    int16
	//    int32
	//    int64
	//    float32
	//    float64
	//    complex64
	//    complex128
	//    bool
	//    string
	//    rune
	//    byte

	var (
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		f32  float32
		f64  float64
		r    rune
		b    byte
		c64  complex64
		c128 complex64
	)

	fmt.Println("printDefaults", i, i8, i16, i32, i64, f32, f64, r, b, c64, c128)
}

func shortDec() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	// EXPECTED OUTPUT
	//  i: 314 f: 3.14 s: Hello b: true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}

// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare
//
//  Declare two variables using multiple short declaration
//
// EXPECTED OUTPUT
//  14 true
// ---------------------------------------------------------

func multShortDec() {
	// ADD YOUR DECLARATIONS HERE
	//
	a, b := 14, true

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(a, b)

}

// ---------------------------------------------------------
// EXERCISE: Short With Expression
//
// 	1. Short declare a variable named `sum`
//
//  2. Initialize it with an expression by adding 27 and 3.5
//
// EXPECTED OUTPUT
//  30.5
// ---------------------------------------------------------

func sumInit() {
	// ADD YOUR DECLARATION HERE
	sum := 27 + 3.5

	// THEN UNCOMMENT THE CODE BELOW
	fmt.Println(sum)
}

// ---------------------------------------------------------
// EXERCISE: Short Discard
//
// 	1. Short declare two bool variables
//     (use multiple short declaration syntax)
//
//  2. Initialize both variables to true
//
//  3. Change your declaration and
//     discard the 2nd variable's value
//     using the blank-identifier
//
//  4. Print only the 1st variable
//
// EXPECTED OUTPUT
//  true
// ---------------------------------------------------------

func shortDiscard() {
	// ADD YOUR DECLARATIONS HERE
	on, _ := true, true

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(on)
}

// ---------------------------------------------------------
// EXERCISE: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 42
//
//     (! You should use redeclaration)
//
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  42, 20, 3.14
// ---------------------------------------------------------

func redeclareFunc() {
	// ADD YOUR DECLARATIONS HERE
	//
	age, yourAge := 100, 20
	ratio, age := 3.14, 42

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(age, yourAge, ratio)
}

// ---------------------------------------------------------
// EXERCISE: Make It Blue
//
//  1. Change `color` variable's value to "blue"
//
//  2. Print it
//
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------

func assignmentFunc() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW:

	color = "blue"
	fmt.Println(color)
}

// ---------------------------------------------------------
// EXERCISE: Variables To Variables
//
//  1. Change the value of `color` variable to "dark green"
//
//  2. Do not assign "dark green" to `color` directly
//
//     Instead, use the `color` variable again
//     while assigning to `color`
//
//  3. Print it
//
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
//
// HINT
//  + operator can concatenate string values
//
//  Example:
//
//  "h" + "e" + "y" returns "hey"
//
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------

func assignmentFunc2() {
	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW

	color = "dark " + color

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	fmt.Println(color)
}


// ---------------------------------------------------------
// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

func multiCaller() {
	// ADD YOUR DECLARATIONS HERE
	_,b := multi()

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}