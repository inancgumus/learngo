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

// ---------------------------------------------------------
// EXERCISE: Adjust the width and height automatically
//
//  In this exercise, your goal is simple:
//
//    Instead of setting the width and height manually,
//    you need to get the width and height of the terminal
//    screen from your operating system.
//
//    Don't worry, it is easier than it sounds. You just
//    need to read a few documentation and install a
//    Go package.
//
//  1. Go here: https://godoc.org/golang.org/x/crypto/ssh/terminal
//
//     Download the package:
//     go get -u golang.org/x/crypto/ssh/terminal
//
//  2. Find the function that gives you the width and height
//     of the terminal.
//
//  3. Call that function from your program and get the
//     width and height.
//
//  4. When an error occurs while retrieving the width
//     and height, report it.
//
//  5. Set the width and height of the board.
//
//  6. After solving the above steps, update your program
//     to use my screen package instead. It offers an
//     easier way to get the width and height of a
//     terminal.
//
//     go get -u https://github.com/inancgumus/screen
//
//
// BONUS
//
//  1. When you set the width, you may see that the ball
//     goes beyond the left and right borders. This happens
//     because the ball emoji spans to multiple console
//     columns (or cells). Ordinary characters have a
//     single column.
//
//     1. Get the width of the ball emoji using a function
//        from the following package:
//
//        go get -u github.com/mattn/go-runewidth
//
//     2. Divide the width using the rune width of the
//        ball emoji.
//
//  2. Your terminal may have borders, so reduce the
//     height by taking into account the height of
//     your terminal borders.
//
//
// EXPECTED OUTPUT
//
//  When you run the program, the ball should start
//  animating across the total width and height of your
//  terminal screen dynamically.
//
//  Currently you set width and height manually, so it
//  wasn't matter whether your terminal was bigger or
//  smaller, but now it will be!
//
//
// HINT
//
//  Please take a look at this if you get stuck.
//
//  You need to pass the Standard Out file handler
//  to the function that returns you the dimensions.
//
//  Check out my screen package to find out how I'm
//  passing the Standard Out file handler.
//
//     https://github.com/inancgumus/screen/blob/master/dimensions.go
//
// ---------------------------------------------------------

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		// drawing buffer length
		//
		// *2 for extra spaces
		// +1 for newlines
		bufLen = (width*2 + 1) * height
	)

	var (
		px, py int    // ball position
		vx, vy = 1, 1 // velocities

		cell rune // current cell (for caching)
	)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// create a drawing buffer
	buf := make([]rune, 0, bufLen)

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

		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		// put the new ball
		board[px][py] = true

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
