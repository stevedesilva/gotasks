package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	digitalClock()
}

// go run *
func digitalClock() {

	screen.Clear()
	for {

		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		ssec := now.Nanosecond() / 1e8

		// fmt.Printf("hour %d, min %d, sec %d\n", hour, min, sec)
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[ssec],
		}

		// alarmed := sec%10 == 0

		// 0 to 4
		for line := range clock[0] {

			// if alarmed {
			// 	// replace clock array with alarm array
			// 	clock = alarm
			// }

			// draw clock array line by line (e.g. horizonally old 8bit computer game style)
			for index, digit := range clock {

				// e.g. print [0][0], [1][0], [2][0]...
				next := clock[index][line]
				// colon blink
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")

			}
			fmt.Println()
		}
		// fmt.Println()
		// time.Sleep(time.Second)

		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}

}
