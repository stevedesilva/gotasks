// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// e1()
	// e2()
	// e3()
	// e1Append()
	// e2Append()
	// e3Append()
	// e3Sort()
	// e4HousePrices()
	// e5Slice()
	// slicingByArguments()
	// sliceHousePrices()
	// test()
	prettyCap()
}

// ---------------------------------------------------------
// EXERCISE: Assign slice literals
//
//   1. Assign the following data using slice literals to the slices that
//      you've declared in the first exercise.
//
//    1. The names of your best three friends (to the names slice)
//
//    2. The distances to five different locations (to the distances slice)
//
//    3. Five bytes of data (to the data slice)
//
//    4. Two currency exchange ratios (to the ratios slice)
//
//    5. Up/Down status of four different web servers (to the alives slice)
//
//  2. Print their type, length and whether they're equal to nil value or not
//
//  3. Compare the length of the distances and the data slices; print a message
//     if they are equal (use an if statement).
//
//
// EXPECTED OUTPUT
//  names    : []string 3 false
//  distances: []int 5 false
//  data     : []uint8 5 false
//  ratios   : []float64 2 false
//  alives   : []bool 4 false
//  The length of the distances and the data slices are the same.
// ---------------------------------------------------------
func e1() {

	var (
		names     []string  = []string{"Max", "Dave", "Paul"}
		distances []int     = []int{3, 4, 5, 6, 7}
		data      []uint8   = []uint8{'A', 'B', 'N', 128, 5}
		ratios    []float64 = []float64{2.4, 4.5}
		alives    []bool    = []bool{true, true, false, false}
	)

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)

}

// ---------------------------------------------------------
// EXERCISE: Declare the arrays as slices
//
//   1. First run the following program as it is
//
//   2. Then change the array declarations to slice declarations
//
//   3. Observe whether anything changes or not (on the surface :)).
//
// EXPECTED OUTPUT
//  names    : []string ["Einstein" "Tesla" "Shepard"]
//  distances: []int [50 40 75 30 125]
//  data     : []uint8 [72 69 76 76 79]
//  ratios   : []float64 [3.14]
//  alives   : []bool [true false true false]
//  zero     : []uint8 []
// ---------------------------------------------------------
// pwd
func e2() {
	names := []string{"Einstein", "Tesla", "Shepard"}
	distances := []int{50, 40, 75, 30, 125}
	data := []byte{'H', 'E', 'L', 'L', 'O'}
	ratios := []float64{3.14145}
	alives := []bool{true, false, true, false}
	zero := []byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)
}

// ---------------------------------------------------------
// EXERCISE: Fix the problems
//
//  1. Uncomment the code
//
//  2. Fix the problems
//
//  3. BONUS: Simplify the code
//
//
// EXPECTED OUTPUT
//   "Einstein and Shepard and Tesla"
//   ["Fire" "Kafka's Revenge" "Stay Golden"]
//   [1 2 3 5 6 7 8 9]
// ---------------------------------------------------------

func e3() {

	names := []string{"Einstein", "Shepard", "Tesla"}

	books := []string{"Stay Golden", "Fire", "Kafka's Revenge"}
	sort.Strings(books)

	// -----------------------------------
	// this time, do not change the nums array to a slice
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums[:])

	// -----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	fmt.Printf("%q\n", strings.Join(names, " and "))

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------
func e1Append() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, png...)
	// 3. compare the slices using the bytes.Equal function
	equal := bytes.Equal(png, header)
	// 4. print whether they're equal or not
	fmt.Println("Is equal: ", equal)

}

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

func e2Append() {
	var (
		pizza       []string
		departures  []time.Time
		graduations []int
		lights      []bool
	)

	pizza = append(pizza, "chicken", "onion", "cheese")
	departures = append(departures, time.Now(), time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*48))
	graduations = append(graduations, 1999, 2000, 2010, 2019)
	lights = append(lights, true, false)

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\ngraduations : %d\n", graduations)
	fmt.Printf("\nlights      : %t\n", lights)

}

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func e3Append() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("toppings       : %s\n", pizza)
}

// ---------------------------------------------------------
// EXERCISE: Append and Sort Numbers
//
//  We'll have a []string that should contain numbers.
//
//  Your task is to convert the []string to an int slice.
//
//  1. Get the numbers from the command-line
//
//  2. Append them to an []int
//
//  3. Sort the numbers
//
//  4. Print them
//
//  5. Handle the error cases
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    provide a few numbers
//
//  go run main.go 4 6 1 3 0 9 2
//    [0 1 2 3 4 6 9]
//
//  go run main.go a b c
//    []
//
//  go run main.go 4 a 1 c d 9
//    [1 4 9]
//
// ---------------------------------------------------------

func e3Sort() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a few numbers.")
		return
	}
	var in []int
	for _, v := range args {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		in = append(in, i)
	}

	sort.Ints(in)
	fmt.Printf("%d", in)
}

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func e5Slice() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	numStr := strings.Fields(data)

	var num, odd, even []int
	for _, v := range numStr {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		num = append(num, i)
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	middle := num[2:4]
	firstTwo := num[:2]
	lastTwo := num[len(num)-2:]
	evenlast1 := even[len(even)-1:]
	oddlast2 := odd[len(odd)-2:]

	fmt.Printf("nums         :%d \n", num)
	fmt.Printf("even         :%d \n", even)
	fmt.Printf("odd          :%d \n", odd)
	fmt.Printf("middle       :%d \n", middle)
	fmt.Printf("first 2      :%d \n", firstTwo)
	fmt.Printf("last 2       :%d \n", lastTwo)
	fmt.Printf("evens last 1 :%d \n", evenlast1)
	fmt.Printf("odd last 2   :%d \n", oddlast2)

}

// ---------------------------------------------------------
// EXERCISE: Slicing by arguments
//
//   We've a []string, you will get arguments from the command-line,
//   then you will print only the elements that are requested.
//
//   1. Print the []string (it's in the code below)
//
//   2. Get the starting and stopping positions from the command-line
//
//   3. Print the []string slice by slicing it using the starting and stopping
//      positions
//
//   4. Handle the error cases (see the expected output below)
//
//   5. Add new elements to the []string slice literal.
//      Your program should work without changing the rest of the code.
//
//   6. Now, play with your program, get a deeper sense of how the slicing
//      works.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: we expect only two arguments)
//
//  go run main.go 1 2 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: starting index >= 0 && stopping pos <= len(slice) )
//
//  go run main.go -1 5
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  (error because: starting <= stopping)
//
//  go run main.go 3 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  go run main.go 0
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus Warsaw]
//
//
//  go run main.go 1
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus Warsaw]
//
//
//  go run main.go 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Nexus Warsaw]
//
//
//  go run main.go 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Warsaw]
//
//
//  go run main.go 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    []
//
//
//  go run main.go 0 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus]
//
//
//  go run main.go 1 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus]
//
//
//  go run main.go 1 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan]
//
// ---------------------------------------------------------

func slicingByArgumentsOld() {

	const (
		msg = `["Normandy" "Verrikan" "Nexus" "Warsaw"]
	
Provide only the [starting] and [stopping] positions
		`
		wrongMsg = `["Normandy" "Verrikan" "Nexus" "Warsaw"]
	
Wrong positions
		`

		convMsg = "Can't convert"
	)

	var (
		starting, stopping, l int
	)
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	args := os.Args[1:]

	if l = len(args); l > 2 || l <= 0 {
		fmt.Println(msg)
		return
	}

	starting, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(convMsg)
		return
	}

	if l == 1 {
		stopping = len(ships)
	} else {
		end, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(convMsg)
			return
		}
		stopping = end
	}

	// starting index >= 0 && stopping pos <= len(slice)
	if validRange, validPos := starting >= 0 && stopping <= len(ships), starting <= stopping; !validRange || !validPos {
		fmt.Println(wrongMsg)
	}

	slice := ships[starting:stopping]
	fmt.Println("Cap ", cap(ships))
	fmt.Println(slice)
}

func slicingByArguments() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])
	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])
}

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func e4HousePrices() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		location                                      []string
		size, beds, baths, prices                     []int
		sizeTotal, bedsTotal, bathsTotal, pricesTotal int
	)
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		col := strings.Split(row, separator)
		location = append(location, col[0])

		s, err := strconv.Atoi(col[1])
		if err != nil {
			return
		}
		sizeTotal += s
		size = append(size, s)

		b, err := strconv.Atoi(col[2])
		if err != nil {
			return
		}
		bedsTotal += b
		beds = append(beds, b)

		ba, err := strconv.Atoi(col[3])
		if err != nil {
			return
		}
		bathsTotal += ba
		baths = append(baths, ba)

		p, err := strconv.Atoi(col[4])
		if err != nil {
			return
		}
		pricesTotal += p
		prices = append(prices, p)

	}

	hdr := strings.Split(header, separator)
	for _, h := range hdr {
		fmt.Printf("%-15s", h)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))

	for i := range rows {
		fmt.Printf("%-15s", location[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-15s", "")
	fmt.Printf("%-15.2f", float64(sizeTotal)/float64(len(size)))
	fmt.Printf("%-15.2f", float64(bedsTotal)/float64(len(beds)))
	fmt.Printf("%-15.2f", float64(bathsTotal)/float64(len(baths)))
	fmt.Printf("%-15.2f", float64(pricesTotal)/float64(len(prices)))
	fmt.Println()
}

// ---------------------------------------------------------
// EXERCISE: Slicing the Housing Prices
//
//  We have received housing prices. Your task is to print only the requested
//  columns of data (see the expected output).
//
//  1. Separate the data by the newline ("\n") -> rows
//
//  2. Separate each row of the data by the separator (",") -> columns
//
//  3. Find the from and to positions inside the columns depending
//     on the command-line arguments.
//
//  4. Print only the requested column headers and their data
//
//
// RESTRICTIONS
//
//  + You should use slicing when printing the columns.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  go run main.go Location
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  NOTE : Finds the position of the Size column and slices the columns
//         appropriately.
//
//  go run main.go Size
//    Size           Beds           Baths          Price
//
//    100            2              1              100000
//    150            3              2              200000
//    200            4              3              400000
//    500            10             5              1000000
//
//
//  NOTE : Finds the positions of the Size and Baths columns and
//         slices the columns appropriately.
//
//  go run main.go Size Baths
//    Size           Beds           Baths
//
//    100            2              1
//    150            3              2
//    200            4              3
//    500            10             5
//
//
//  go run main.go Beds Price
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  Note : It works even if the given column name does not exist.
//
//  go run main.go Beds NotExist
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  go run main.go NotExist NotExist
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
// Note : It works even if the Price's index > Size's index
//
//        In that case, it resets the invalid starting position to 0.
//
//        But it still uses the Size column.
//
//  go run main.go Price Size
//    Location       Size
//
//    New York       100
//    New York       150
//    Paris          200
//    Istanbul       500
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func sliceHousePrices() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// parse the data
	rows := strings.Split(data, "\n")
	cols := strings.Split(rows[0], separator)

	// default case: slice for all the columns
	from, to := 0, len(cols)

	// find the from and to positions depending on the command-line arguments
	args := os.Args[1:]
	for i, v := range cols {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}

		if l == 2 && v == args[1] {
			to = i + 1 // +1 because the stopping index is a position
		}
	}

	// from cannot be greater than to: reset invalid arg to 0
	if from > to {
		from = 0
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		// print only the requested columns
		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		// print extra new line for the header
		if i == 0 {
			fmt.Println()
		}
	}
}

func test() {
	arr := [...]int{9, 7, 5, 3, 1}
	fmt.Println("arr: ", arr)

	nums := arr[2:]
	fmt.Println("nums: ", nums)
	nums2 := nums[1:]
	fmt.Println("nums2: ", nums2)

	fmt.Println(strings.Repeat("=", 50))

	arr[2]++
	fmt.Println("arr[2]++: ", arr)
	fmt.Println()

	// x := 6 - -1 - 4
	// fmt.Println(x)
	fmt.Println("Before: ", nums[1])
	fmt.Println("nums[1]: ", nums[1])
	fmt.Println("arr[4]: ", arr[4])
	fmt.Println("arr[4] - 4: ", arr[4]-4)
	fmt.Println("3 - 1 - 4: ", arr[4]-4)

	// 3 - (1 - 4)
	// nums[1] = nums[1] - (arr[4] - 4)
	nums[1] -= arr[4] - 4
	fmt.Println("After:  ", nums[1])

	// fmt.Println("nums[1] :",nums[1])
	// fmt.Println("arr[4]  :",arr[4])

	// // fmt.Println("arr[4] - 4 :",arr[4] - 4)
	// fmt.Println("nums[1] -= arr[4] - 4: ",nums)

	nums2[1] += 5
	// fmt.Println("nums2[1] += 5: ",nums2)
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println(nums)
}

func prettyCap() {
	s.PrintBacking = true

	var nums []int
	s.Show("no backing array", nums)

	nums = append(nums, 1, 3)
	s.Show("allocates", nums)

	nums = append(nums, 2)
	s.Show("free capacity", nums)

	nums = append(nums, 4)
	s.Show("no allocation", nums)

	nums = append(nums, nums[2:]...)
	s.Show("nums <- nums[2:]", nums)

	nums = append(nums[:2], 7, 9)
	s.Show("nums[:2] <- 7, 9", nums)


	nums = nums[:6]
	s.Show("nums[:6]", nums)
	// backing := []int{1, 2, 3, 4, 5, 6}
	// s.Show("backing", backing)

	// b := backing

	// b = b[:1]
	// s.Show("b", b)

	// b = b[:cap(b)]
	// s.Show("b", b)

}
