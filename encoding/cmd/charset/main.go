package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ascii - go run main.go 32 127
// emoji - go run main.go 128512 128591
func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if stop == 0 || start == 0 {
		stop, start = 'Z', 'A'
	}
	fmt.Printf("%-10s %-10s %-10s %-10s\n%s\n", "literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	for n := start; n <= stop; n++ {
		// space denotes utf8 encoding % -12[1]x\n
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}
