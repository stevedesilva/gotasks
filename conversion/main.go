package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------

func main() {
	f1()
	f2()
	f3()
	f4()
	f5()
}

func f1() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

// EXPECTED OUTPUT
//  10.5
func f2() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

func f3() {
	fmt.Println(float64(5.5))
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

func f4() {
	age := 2
	fmt.Println(7.5 + float64(age))
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func f5() {
	// DO NOT TOUCH THIS VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}