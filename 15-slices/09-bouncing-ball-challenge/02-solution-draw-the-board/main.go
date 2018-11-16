// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

const (
	width  = 25
	height = 8
)

func main() {
	// -------------------------------------------------
	// CREATE THE BOARD
	// -------------------------------------------------
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

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
	fmt.Print(string(buf))
}
