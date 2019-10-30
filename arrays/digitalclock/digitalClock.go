package main

import (
	"fmt"
	"time"
)

func main() {
	digitalClock()
}
func digitalClock() {
	// print digits

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	var digits = [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	fmt.Printf("hour %d, min %d, sec %d\n", hour, min, sec)
	clock := [...]placeholder{
		digits[hour/10],digits[hour%10],
		colon,
		digits[min/10],digits[min%10],
		colon,
		digits[sec/10],digits[sec%10],
	}

	// for _, digit := range digits {
	// 	for _, line := range digit {
	// 		fmt.Println(line)
	// 	}
	// 	fmt.Println()
	// }

	// 0 to 4
	for line := range clock[0] {
		// get each digit
		for digit := range clock {
			// print [0][0], [1][0], [2][0]...
			fmt.Print(clock[digit][line], " ")
		}
		fmt.Println()
	}

}
