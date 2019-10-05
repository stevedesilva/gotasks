package example

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	// feetInMeters float64 = 0.3048
	// feetInYards          = feetInMeters / 0.9144
	feetInMeters = 0.3048
	feetInYards  = feetInMeters / 0.9144
)

// Conversion func
func Conversion() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect args length.")
		return
	}
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("error: %v is not a number.\n", feet)
		return
	}
	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%T\n", feetInYards)
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
