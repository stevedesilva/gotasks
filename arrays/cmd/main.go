package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// emptyArray()
	// getAndSetArrayElements()
	// arrayLiterals()
	// simplify()
	// comparableArrays()
	// assignmentOfArrays()
	// moodlySingleArray()
	// moodlyMultiArray()
	// wizardPrinter()
	// currencyConverter()
	// searchEngine()
	// average()
	// numberSorter()
	wordSearch()
}

// ---------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your best three friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// HINT
//   When printing the elements of an array, you can use the usual Printf verbs.
//
//   For example:
//     When printing a string array, you can use "%q" verb as usual.
//
// EXPECTED OUTPUT
//  names    : [3]string{"", "", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []
// ---------------------------------------------------------
func emptyArray() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)
	// 1. Declare and print the arrays with their types.
	fmt.Printf("names    : %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("alives   : %#v\n", alives)
	fmt.Printf("zero     : %#v\n", zero)

	//  2. Print only the types of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("alives   : %T\n", alives)
	fmt.Printf("zero     : %T\n", zero)

	//  3. Print only the elements of the same arrays.
	fmt.Println()
	fmt.Printf("names    : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratios   : %.2f\n", ratios)
	fmt.Printf("alives   : %t\n", alives)
	fmt.Printf("zero     : %d\n", zero)

}

// ---------------------------------------------------------
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func getAndSetArrayElements() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	names[0], names[1], names[2] = "Tyson", "Ali", "Frazier"

	distances[0] = 50
	distances[1] = 40
	distances[2] = 75
	distances[3] = 30
	distances[4] = 125

	data[0] = 'H'
	data[1] = 'E'
	data[2] = 'L'
	data[3] = 'L'
	data[4] = 'O'

	ratios[0] = 3.14145

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %c\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	fmt.Printf(`
%s
FOR RANGES
%[1]s
`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}

}

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func arrayLiterals() {

	names := [...]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [...]uint8{'H', 'E', 'L', 'L', 'O'}
	ratios := [...]float64{3.14145}
	alives := [...]bool{true, false, true, false}
	var zero [0]uint8

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Printf(`
%s
FOR RANGES
%[1]s
`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}

}

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------
func simplify() {
	var names = [...]string{"Einstein", "Shepard", "Tesla"}
	var books = [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

// ---------------------------------------------------------
// EXERCISE: Compare the Arrays
//
//  1. Uncomment the code
//
//  2. Fix the problems so that arrays become comparable
//
// EXPECTED OUTPUT
//  true
//  true
//  false
// ---------------------------------------------------------
func comparableArrays() {
	week := [...]string{"Monday", "Tuesday"}
	wend := [...]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]uint8

	fmt.Println(data == image)
}

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------

func assignmentOfArrays() {
	books := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}
	var (
		upper, lower [3]string
	)

	for i, v := range books {
		upper[i] = strings.ToUpper(v)
		lower[i] = strings.ToLower(v)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------
func moodlySingleArray() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}

	name := args[0]
	var moods = [...]string{
		"good 👍", "happy 😀", "awesome 😎",
		"terrible 😩", "sad 😞", "bad 👎",
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(moods))
	fmt.Printf("%s feels %s\n", name, moods[random])

}

// ---------------------------------------------------------
// EXERCISE: Moodly #2
//
//   This challenge is based on the previous Moodly challenge.
//
//   1. Display the usage if the username or the mood is missing
//
//   2. Change the moods array to a multi-dimensional array.
//
//      So, create two inner arrays:
//        1. One for positive moods
//        2. Another one for negative moods
//
//   4. Randomly select and print one of the mood messages depending
//      on the given mood command-line argument.
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name] [positive|negative]
//
//   go run main.go Socrates
//     [your name] [positive|negative]
//
//   go run main.go Socrates positive
//     Socrates feels good 👍
//
//   go run main.go Socrates positive
//     Socrates feels happy 😀
//
//   go run main.go Socrates positive
//     Socrates feels awesome 😎
//
//   go run main.go Socrates negative
//     Socrates feels bad 👎
//
//   go run main.go Socrates negative
//     Socrates feels sad 😞
//
//   go run main.go Socrates negative
//     Socrates feels terrible 😩
// ---------------------------------------------------------
func moodlyMultiArray() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var moodVal int
	if mood != "positive" {
		moodVal = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[moodVal][n])
}

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func wizardPrinter() {
	table := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}
	fmt.Printf("%-15s %-15s %-15s\n", "First Name", "Last Name", "Nickname")
	fmt.Println(strings.Repeat("=", 50))
	for _, v := range table {
		fmt.Printf("%-15s %-15s %-15s\n", v[0], v[1], v[2])
	}
}

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------
func currencyConverter() {

	const (
		EUR = iota
		GBP
		JPY
	)

	ratio := [3]float64{0.90, 0.77, 108.67}

	in := os.Args[1:]
	if len(in) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	n, err := strconv.ParseFloat(in[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", n, n*ratio[EUR])
	fmt.Printf("%.2f USD is %.2f GBP\n", n, n*ratio[GBP])
	fmt.Printf("%.2f USD is %.2f JPY\n", n, n*ratio[JPY])

}

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

func searchEngine() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("    Tell me a book title.")
		return
	}
	word := strings.ToLower(args[0])
	var flag bool

	fmt.Println("Search Results:")
	for _, v := range books {
		if strings.Contains(strings.ToLower(v), word) {
			fmt.Println("+", v)
			flag = true
		}
	}
	if !flag {
		fmt.Printf("We don't have the book: %q\n", word)
	}

}

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func average() {
	args := os.Args[1:]

	if len := len(args); len < 1 || len > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		nums  [5]float64
		sum   float64
		count int
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		count++
		nums[i] = n
		sum += n
	}

	fmt.Printf("Your numbers: %v\n", nums)
	fmt.Printf("Average: %v \n", sum/float64(count))

}

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func numberSorter() {
	args := os.Args[1:]

	if l := len(args); l == 0 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	} else if l > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	var (
		nums [5]float64
	)

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		nums[i] = n
	}

	fmt.Printf("Received numbers: %v\n", nums)
	sorted := bubblesort(nums)
	fmt.Printf("Your numbers: %v\n", sorted)

}

func bubblesort(nums [5]float64) [5]float64 {
	/*
		check whether it's the last element or not:
		  i < len(nums)-1
		check whether the next number is greater than the current one, if so, swap it:
		  v > nums[i+1]
	*/
	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums

}

// ---------------------------------------------------------
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func wordSearch() {

	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	words := strings.Fields(strings.ToLower(corpus))

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, v := range filter {
			if q == v {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}


