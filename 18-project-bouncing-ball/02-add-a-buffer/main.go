// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

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

	board[0][0] = true
	board[9][0] = true
	board[9][49] = true
	board[0][49] = true

	buffer := make([]rune, 0, (width*2+1)*height)

	for _, row := range board {
		for y := range row {
			if row[y] == true {
				cell = ball
			} else {
				cell = empty
			}
			buffer = append(buffer, cell, ' ')
		}
		buffer = append(buffer, '\n')
	}
	fmt.Println(string(buffer))
}
