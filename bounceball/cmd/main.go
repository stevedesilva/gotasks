package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
)
func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = false  // prints the backing arrays
	s.MaxPerLine = 15       // prints max 15 elements per line
	// s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}

func main() {
	const (
		width     = 50
		height    = 10
		cellBall  = '⚾'
		cellEmpty = ' '
	)

	var cell rune
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	board[12][2] = true
	board[16][2] = true
	board[14][4] = true
	board[10][6] = true
	board[18][6] = true
	board[12][7] = true
	board[14][7] = true
	board[16][7] = true

	buf := make([]rune, 0, width*height)
	s.Show("buf start",buf)
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		// s.Show("buf" + string(i),buf)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		fmt.Print(string(buf))
		s.Show("buf " + string(i),buf)
	}
}
