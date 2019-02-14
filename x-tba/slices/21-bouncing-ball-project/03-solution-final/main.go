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

	"github.com/inancgumus/screen"
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

	var (
		X, Y       int    // ball positions
		xVel, yVel = 1, 1 // velocities
	)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// -------------------------------------------------
		// CALCULATE THE NEXT BALL POSITION
		// -------------------------------------------------
		X += xVel
		Y += yVel

		// when the ball hits the borders change its direction
		// by changing its velocity
		if X <= 0 || X >= width-1 {
			xVel *= -1
		}
		if Y <= 0 || Y >= height-1 {
			yVel *= -1
		}

		// -------------------------------------------------
		// CLEAR THE BOARD AND PUT THE BALL
		// -------------------------------------------------
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}
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
					ball = '🏈'
				}
				buf = append(buf, ball, ' ')
			}
			buf = append(buf, '\n')
		}

		// clear the screen and draw the board
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		// TODO: in the notes this is at the beginning of the loop

		// draw after a while: slows down the animation
		time.Sleep(time.Second / 20)
	}
}
