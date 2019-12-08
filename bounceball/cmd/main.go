package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// time go run main.go > /dev/null
func main() {
	const (
		width     = 50
		height    = 10
		cellBall  = '⚾'
		cellEmpty = ' '
		maxframe  = 1200
		speed     = time.Second / 20
	)

	var (
		cell   rune
		px, py int
		vx, vy = 1, 1
	)
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxframe; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

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
		screen.MoveTopLeft()

		fmt.Print(string(buf))
		time.Sleep(speed)
	}

}
