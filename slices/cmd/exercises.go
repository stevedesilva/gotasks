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
	"sort"
	"strings"
	"time"
)

func main() {
	// e1()
	// e2()
	// e3()
	// e1Append()
	// e2Append()
	e3Append()
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
	pizza = append(pizza, "onions","extra cheese")

	fmt.Printf("toppings       : %s\n", pizza)
}
