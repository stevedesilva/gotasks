// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// e1()
	// e2()
	e3()
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
