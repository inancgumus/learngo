// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	// get the width and height of the terminal dynamically ***
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("cannot get width and height", err)
		return
	}
	width /= 2 // our emoji is 2 chars wide

	var (
		px, py   int    // ball position
		ppx, ppy int    // previous ball position ***
		vx, vy   = 1, 1 // velocities

		cell rune // current cell (for caching)
	)

	// create the board
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// check whether the ball goes beyond the borders
		if !(px >= width || py >= height) {
			// remove the previous ball and put the new ball
			board[px][py], board[ppx][ppy] = true, false

			// save the previous positions
			ppx, ppy = px, py
		}

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
