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

const (
	width  int = 50
	height int = 10

	empty rune = ' '
	ball       = '⚾'
)

func main() {
	var cell rune

	board := make([][]bool, height)

	for row := range board {
		board[row] = make([]bool, width)
	}

	board[9][49] = true

	for _, row := range board {
		for y := range row {
			if row[y] == true {
				cell = ball
			} else {
				cell = empty
			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}
