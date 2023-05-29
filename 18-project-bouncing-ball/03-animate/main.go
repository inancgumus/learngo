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
	"time"

	"github.com/inancgumus/screen"
)

const (
	width  int = 50
	height int = 10

	empty rune = ' '
	ball       = '⚾'

	maxFrames = 1000
	speed     = time.Second / 20
)

func main() {
	var (
		cell   rune
		px, py int
		vx, vy = 1, 1
	)

	board := make([][]bool, height)

	for row := range board {
		board[row] = make([]bool, width)
	}

	buffer := make([]rune, 0, (width*2+1)*height)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px == 0 || px == height-1 {
			vx *= -1
		}

		if py == 0 || py == width-1 {
			vy *= -1
		}

		board[px][py] = true
		buffer = buffer[:0]
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
		screen.MoveTopLeft()
		fmt.Println(string(buffer))
		board[px][py] = false

		time.Sleep(speed)
	}
}
