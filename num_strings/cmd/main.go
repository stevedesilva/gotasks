package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Do Some Calculations
//
//  1. Print the sum of 50 and 25
//  2. Print the difference of 50 and 15.5
//  3. Print the product of 50 and 0.5
//  4. Print the quotient of 50 and 0.5
//  5. Print the remainder of 25 and 3
//  6. Print the negation of `5 + 2`
//
// EXPECTED OUTPUT
//  75
//  34.5
//  25
//  100
//  1
//  -7
// ---------------------------------------------------------

func q1() {
	sum := 50 + 25
	fmt.Println("sum: ", sum)

	diff := float64(50) - 15.5
	fmt.Println("diff: ", diff)

	product := float64(50) * .5
	fmt.Println("product: ", product)

	quotient := float64(50) / .5
	fmt.Println("quotient: ", quotient)

	negation := 25 % 3
	fmt.Println("negation: ", negation)

	remainder := -(5 + 2)
	fmt.Println("remainder: ", remainder)
}

// ---------------------------------------------------------
// EXERCISE: Fix the Float
//
//  Fix the program to print 2.5 instead of 2
//
// EXPECTED OUTPUT
//  2.5
// ---------------------------------------------------------
func q2() {
	x := 5 / float64(2)
	fmt.Println(x)
}

// ---------------------------------------------------------
// EXERCISE: Precedence
//
//  Change the expressions to produce the expected outputs
//
// RESTRICTION
//  Use parentheses to change the precedence
// ---------------------------------------------------------

func question3() {
	// This expression should print 20
	print20()

	// This expression should print -16
	shouldMinusPrint16()

	// This expression should print -25
	printMinus25()

	// This expression should print 0.5
	print0Point5()

	// This expression should print 24
	print24()

	// This expression should print 15
	print15()

	// This expression should print 40
	// Note that you may need to use floats to solve this
	print40()
}

func print20() {
	fmt.Println(10 + 5 - (5 - 10))
}

// 100 / 5 / 2
func print40() {
	sum := 100 / (5. / 2.)
	fmt.Println(sum)

}

// 10 / 2 * 10 % 7
func print15() {
	sum := 10 / 2 * (10 % 7)
	fmt.Println(sum)
}

// 3 + 1/2*10 + 4
func print24() {
	fmt.Println((3+1)/2*10 + 4)
}

func print0Point5() {
	fmt.Println(0.5 * (2 - 1))
}

// 5 + 10*2 - 5
func printMinus25() {
	fmt.Println(5 + 10*(2-5))
}

// -10 + 0.5 - 1 + 5.5
func shouldMinusPrint16() {
	fmt.Println(-10 + 0.5 - (1 + 5.5))
}

// ---------------------------------------------------------
// EXERCISE: Incdecs
//
//  1. Increase the `counter` 5 times
//  2. Decrease the `factor` 2 times
//  3. Print the product of counter and factor
//
// RESTRICTION
//  Use only the incdec statements
//
// EXPECTED OUTPUT
//  -75
// ---------------------------------------------------------

func question4() {
	// DO NOT TOUCH THIS
	counter, factor := 45, 0.5

	// TYPE YOUR CODE BELOW
	for i := 0; i < 5; i++ {
		counter++
		fmt.Printf("count value %d \n", counter)
	}

	for i := 0; i < 2; i++ {
		factor--
		fmt.Printf("factor value %g \n", factor)
	}

	sum := float64(counter) * factor

	fmt.Printf("product of counter(%d) and factor(%g) is %g\n", counter, factor, sum)

}

// ---------------------------------------------------------
// EXERCISE: Manipulate a Counter
//
//  1. Write the simplest line of code to increase
//     the counter variable by 1.
//
//  2. Write the simplest line of code to decrease
//     the counter variable by 1.
//
//  3. Write the simplest line of code to increase
//     the counter variable by 5.
//
//  4. Write the simplest line of code to multiply
//     the counter variable by 10.
//
//  5. Write the simplest line of code to divide
//     the counter variable by 5.
//
// EXPECTED OUTPUT
//  10
// ---------------------------------------------------------

func question5() {
	// DO NOT CHANGE THE CODE BELOW
	var counter int

	// TYPE YOUR CODE HERE
	counter++
	counter--

	counter += 5

	counter *= 10

	counter /= 5

	// DO NOT CHANGE THE CODE BELOW
	fmt.Println(counter)
}

// ---------------------------------------------------------
// EXERCISE: Simplify the Assignments
//
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

func question6() {
	width, height := 10, 2

	// width = width + 1
	// width = width + height
	// width = width - 1
	// width = width - height
	// width = width * 20
	// width = width / 25
	// width = width % 5

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)
}

// ---------------------------------------------------------
// EXERCISE: Circle Area
//
//  Calculate the area of a circle from the given radius
//
// CIRCLE AREA FORMULA
//  area = πr²
//  https://en.wikipedia.org/wiki/Area#Circles
//
// HINT
//  For PI you can use `math.Pi`
//
// EXPECTED OUTPUT
//  radius: 10 -> area: 314.1592653589793
//
// BONUS EXERCISE!
//  1. Print the area as 314.16
//  2. To do that you need to use the correct Printf verb :)
//      Instead of `%g` verb below.
//
//    EXPECTED OUTPUT
//     radius: 10 -> area: 314.16
// ---------------------------------------------------------

func question7() {
	var (
		radius = 10.
		area   float64
	)

	// ADD YOUR CODE HERE
	// ...
	pi := math.Pi
	fmt.Printf("Pi is %.2f\nPi is %[1]g\n", pi)
	area2 := pi * (radius * radius)
	_ = area2

	area = math.Pi * math.Pow(radius, 2)

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %g\n", radius, area)
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}

// ---------------------------------------------------------
// EXERCISE: Sphere Area
//
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the surface area of a sphere
//
// SPHERE SURFACE AREA FORMULA
//  area = 4πr²
//  https://en.wikipedia.org/wiki/Sphere#Surface_area
//
// RESTRICTION
//  Use `math.Pow` to calculate the area
//  Read its documentation to see how it works.
//  https://golang.org/pkg/math/#Pow
//
// EXPECTED OUTPUT
//  go run main.go 10
//    1256.64
//
//  go run main.go 54.2
//    36915.47
// ---------------------------------------------------------

func question8() {
	var radius, area float64

	// ADD YOUR CODE HERE
	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	area = 4 * math.Pi * math.Pow(radius, 2)

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}

// ---------------------------------------------------------
// EXERCISE: Sphere Volume
//
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the volume of a sphere
//
// SPHERE VOLUME FORMULA
//  https://en.wikipedia.org/wiki/Sphere#Enclosed_volume
//
// RESTRICTION
//  Use `math.Pow` to calculate the volume
//
// EXPECTED OUTPUT
//  go run main.go 10
//    4188.79
//
//  go run main.go .5
//    0.52
// ---------------------------------------------------------

func question9() {
	var radius, vol float64

	// ADD YOUR CODE HERE
	// ...

	// ADD YOUR CODE HERE
	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	// vol = 4/3* math.Pi * math.Pow(radius,3)
	// vol = .523 * math.Pow(radius, 3)

	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func e1() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	// path := "c:\\program files\\duper super\\fun.txt\n" +
	// 	"c:\\program files\\really\\funny.png"

	path :=
		`c:\\program files\\duper super\\fun.txt\
c:\\program files\\really\\funny.png`
	fmt.Println(path)
}

// ---------------------------------------------------------
// EXERCISE: Print JSON
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func e2() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	// json := "\n" +
	// 	"{\n" +
	// 	"\t\"Items\": [{\n" +
	// 	"\t\t\"Item\": {\n" +
	// 	"\t\t\t\"name\": \"Teddy Bear\"\n" +
	// 	"\t\t}\n" +
	// 	"\t}]\n" +
	// 	"}\n"

	json :=
		`{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
}

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func e3() {
	// uncomment the code below
	name := os.Args[1]

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := "hi " + name + "\n" + "how are you?\n"

	fmt.Println(msg)
}

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
//  1. Change the following program to work with unicode
//     characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func e4() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.
	in := os.Args[1]

	length := utf8.RuneCountInString(in)

	fmt.Println(length)
}

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

func e5() {
	msg := os.Args[1]
	length := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", length)

	fmt.Println(s)
}

// ---------------------------------------------------------
// EXERCISE: ToLowercase
//
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
//
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
//
// INPUT
//  "SHEPARD"
//
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------

func e6() {
	in := os.Args[1]
	fmt.Println(strings.ToLower(in))
}

// ---------------------------------------------------------
// EXERCISE: Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
//
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

func e7() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func e8() {
	// currently it prints 17
	// it should print 5

	name := "inanç           "
	name = strings.TrimRight(name, " ")
	fmt.Println(utf8.RuneCountInString(name))
	var t byte = 1
	_ = t
}

func main() {
	// q1()
	// q2()
	// question3()
	// question4()
	// question5()
	// question6()
	// question7()
	// question8()
	// question9()
	// e1()
	// e2()
	// e3()
	// e4()
	// e5()
	// e6()
	// e7()
	// e8()
	num := 7
	fmt.Printf("%032b = %[1]d", num)
}
