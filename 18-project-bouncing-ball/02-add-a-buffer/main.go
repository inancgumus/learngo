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
)

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'
	)

	var cell rune // current cell (for caching)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	// draw a smiley
	board[12][2] = true
	board[16][2] = true
	board[14][4] = true
	board[10][6] = true
	board[18][6] = true
	board[12][7] = true
	board[14][7] = true
	board[16][7] = true

	// use the loop for measuring the performance difference
	for i := 0; i < 1000; i++ {
		// rewind the buffer so that the program reuses it
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
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

		// print the buffer
		fmt.Print(string(buf))
	}
}
