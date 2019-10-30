// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/
func main() {
	/*
		VERSION #1: Declare variables

		Every type has a zero-value:
		  numeric types => 0
		  bool          => false
		  string        => ""
	*/

	// var banner string
	// var board string
	// var turn int
	// var maxTurns int
	// var won bool

	/*
		VERSION #2: Declare variables parallel (same as above)

		turn and lastPos are int
		won and wrongMove are bool
	*/
	// var banner, board string
	// var turn, maxTurns int
	// var won bool

	/*
		VERSION #3: Declare variables in a group and parallel (same as above)
	*/
	// var (
	//  banner, board  string
	// 	turn, maxTurns int
	// 	won            bool
	// )

	/*
		VERSION #4: Declare variables in a group (same as above)
	*/
	var (
		banner string // tictactoe banner
		board  string // tictactoe board

		turn     int // total valid turns played
		maxTurns int // maximum number of turns

		won bool // is there any winner?

		progress float64 // remaining progress
	)

	/*
		#5: Assignment
	*/
	banner = " TIC~TAC~TOE"
	board = `
/---+---+---\
|   | X |   |
+---+---+---+
|   | O |   |
+---+---+---+
|   |   |   |
\---+---+---/`

	// maxTurns = 9
	// turn = 2

	// multiple assignment
	maxTurns, turn = 9, 2

	// cannot assign int to float
	// progress = 1 - (turn / maxTurns) * 100
	// convert ints to float so that the result will be float
	// literals are typeless: they automatically get converted to the surrounding operands
	progress = (1 - (float64(turn) / float64(maxTurns))) * 100

	fmt.Println(banner)
	fmt.Println(board)

	fmt.Println()
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)

	// This is also valid: the expression is evaluated on the fly.
	// fmt.Printf("Turns left             : %.1f%%\n",
	// 	(1 - (float64(turn) / float64(maxTurns))) * 100)
}
