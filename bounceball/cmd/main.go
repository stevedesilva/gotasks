package main

import "fmt"

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

	buf := make([]rune, 0, width*height)

	board[12][2] = true
	board[16][2] = true
	board[14][4] = true
	board[10][6] = true
	board[18][6] = true
	board[12][7] = true
	board[14][7] = true
	board[16][7] = true

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			// fmt.Print(string(cell), " ")
			buf = append(buf, cell, ' ')
		}
		// fmt.Println()
		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))

}
