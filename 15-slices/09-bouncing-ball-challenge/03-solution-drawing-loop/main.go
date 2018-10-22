// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"
)

const (
	width  = 25
	height = 8

	maxFrames = 1200
)

func main() {
	// -------------------------------------------------
	// CREATE THE BOARD
	// -------------------------------------------------
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	var X, Y int // ball positions

	for i := 0; i < maxFrames; i++ {
		// draw after a while: slows down the animation
		time.Sleep(time.Second / 20)

		// -------------------------------------------------
		// PUT THE BALL
		// -------------------------------------------------
		board[X][Y] = true

		// -------------------------------------------------
		// DRAW THE BOARD
		// -------------------------------------------------
		var (
			buf  = make([]rune, 0, width*height)
			ball rune
		)

		for y := range board[0] {
			for x := range board {
				ball = '🎱'
				if board[x][y] {
					ball = '🎾'
				}
				buf = append(buf, ball, ' ')
			}
			buf = append(buf, '\n')
		}

		// clear the screen and draw the board
		fmt.Print("\033[2J", string(buf))
	}
}
