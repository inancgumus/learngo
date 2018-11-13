// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
// Post your solution to twitter with this hashtag:
// #learngoprogramming
//
// Notify me on twitter by adding my account to your tweet:
// @inancgumus
//
// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
// This document contains what you need to do to create
// the bouncing ball animation.
//
// You're going to learn a great deal of knowledge about
// slices and you'll earn a good experience while doing this.
//
// However, refer to this document only when you get stuck.
// Do not follow it step by step.
// Try to solve the challenge on your own.
//
// This document organized into steps/sections.
// So you can jump to that step/section directly.
//
// Good luck!
// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
// #1 Declare constants here.
// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
// The width and the height of the board.
// You're going to draw your board using these.

func main() {
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// -> Declare ball positions: X and Y
	//    Initialize them to 0s.

	// -> Declare ball's velocity: xVelocity and yVelocity
	//
	//    Velocity means: Speed and Direction
	//    X velocity =  1 // balls moves to the right
	//    X velocity = -1 // balls moves to the left
	//    Y velocity =  1 // balls moves down
	//    Y velocity = -1 // balls moves up
	//
	// -> On each step, add velocities to ball's position.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 CREATE THE BOARD
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	//
	// -> Use the make function to initialize your board.
	//    Remember: You also need to initialize each sub-slice.
	//              (in a for loop)
	//
	// -> You can use [][]bool for your board.
	//
	//    Because, when you set one of the items to true,
	//    then you know that the ball is in that position.
	//
	//    EXAMPLE:
	//    false    false    false   false
	//    false    true -+  false   false
	//    false    false |  false   false
	//                   v
	//           the ball is here
	//           board[1][1] is true

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 DRAWING LOOP
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// {
	//
	// Create a loop.
	//
	// On each step of the loop, you're going to:
	//  -> Clear the board
	//  -> Calculate the next ball position
	//  -> Draw the board with the balls

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 CLEAR THE BOARD
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// -> Set all the board elements to false.
	//    (I mean the sub-slices' elements)

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 CALCULATE THE NEXT BALL POSITION
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// -> Add the velocities to the ball's current position:
	//
	//      X += xVelocity
	//      Y += yVelocity

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 When ball hits the borders change its direction.
	//
	//    -> Multiply the velocity by -1 to change its X direction.
	//    -> Do the same for the Y velocity as well.

	// 👉 Set the ball's position in the board.
	//
	//    -> You will use this information when drawing the board.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 DRAW THE BOARD
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// -> Make a large enough []rune `buffer`.
	//
	// HINT: width * height will give you a large enough buffer.
	// TIP : You could also use string but it would be inefficient.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 FILL YOUR BUFFER:
	//
	// + It's better to use buffers for these kind of things.
	// + It's worst to call the Print functions all the time.
	//
	// 1. Loop for the height of the board.
	// 2. Then in a nested loop, loop for the width of the board.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 NESTEP LOOP (WIDTH LOOP):
	//
	// In each step of the nested loop, do this:
	//
	// 1. Check whether the ball is in the x, y positions.
	//    You need to check for it using your board slice.
	//
	// 2. If so, append this tennis ball '🎾' to the buf slice.
	// 3. If not, append this pool ball '🎱' to the buf slice.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 HEIGHT LOOP:
	//
	// After the nested loop (but in the height loop):
	//
	// 1. Append the newline character to the buf: '\n'
	//    This will allow you to print the next row to the
	//    next line.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 🎾 PRINT THE BOARD
	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// After the loop print this to clear the screen.
	//
	//   fmt.Print("\033[2J")
	//
	// Note       : This will only work in Linux and Mac.
	// For Windows: Just install Ubuntu bash on Windows, it's easy now!
	//              It isn't a virtual machine.
	// https://docs.microsoft.com/en-us/windows/wsl/install-win10

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 PRINT YOUR BOARD (USING THE BUFFER):
	//
	// -> Do not forget converting it to string.
	//    Because your buffer is []rune.
	//
	//      fmt.Print(string(buf))
	//
	// You'll learn the details about rune and strings later.

	// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
	// 👉 SLOW DOWN THE SPEED
	// And lastly, call the time.Sleep function to slow down
	// the speed of the loop, so you can see the ball :)
	//
	//   time.Sleep(time.Millisecond * 60)

	// } DRAWING LOOP ENDS HERE 👈
}
