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
	"github.com/mattn/go-runewidth"
)

// ---------------------------------------------------------
// EXERCISE: No Slice
//
//   Can you modify the program so that it doesn't use a
//   slice for the board. You can use a slice for the buffer
//   though.
//
//   In this exercise, you'll understand that you don't
//   have to use slices in any problem you encounter with.
//
//   See what it feels like not using a slice for this
//   solution.
//
//   Think about why you have to use a slice for the buffer?
//
//   Can there be any other solution?
//
// ---------------------------------------------------------

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		ivx, ivy = 1, 1 // initial velocities
	)

	var (
		px, py   int        // ball position
		ppx, ppy int        // previous ball position
		vx, vy   = ivx, ivy // velocities

		cell rune // current cell (for caching)
	)

	width, height := screen.Size()
	width /= runewidth.RuneWidth(cellBall)
	height-- // there is a 1 pixel border in my terminal

	// REMOVE THIS: create a single-dimensional board
	board := make([]bool, width*height)

	// create a drawing buffer
	// *2 for extra spaces
	// +1 for newlines
	buf := make([]rune, 0, (width*2+1)*height)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		if py <= 0 || py >= height-ivy {
			vy *= -1
		}

		// remove the previous ball and put the new ball
		pos := py*width + px
		ppos := ppy*width + ppx
		ppx, ppy = px, py

		board[pos], board[ppos] = true, false

		buf = buf[:0]

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cell = cellEmpty

				if board[y*width+x] {
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
