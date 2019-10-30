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
	var (
		won  bool
		turn int
	)

	// VERSION 1
	// const banner = " TIC~TAC~TOE"
	// const board = `
	// /---+---+---\
	// | X | O | X |
	// +---+---+---+
	// | X | X |   |
	// +---+---+---+
	// | O | O | O |
	// \---+---+---/`
	//
	// const maxTurns = 9

	// VERSION 2
	// Benefit of constants:
	// + Resolved to literals at compile-time
	// + Fast at runtime
	// + Provides overwriting unlike variables
	// + Typeless: Can change type
	const (
		banner = `
 TIC~TAC~TOE

/---+---+---\
| X | O | X |
+---+---+---+
| X | X |   |
+---+---+---+
| O | O | O |
\---+---+---/`

		maxTurns = 9
	)

	// Constants have default types
	// 1    => int
	// 1.5  => float64
	// true => bool
	// "hi" => string
	// 'a'  => rune

	// You cannot assign to constants
	// banner = "TIC TAC TOE"
	// maxTurns = 10

	// no need: float64(maxTurns) — constants are typeless (like basic literals)
	// if maxTurns was `const maxTurns int`; then it'd be needed.
	progress := (1 - (float64(turn) / maxTurns)) * 100

	fmt.Println(banner)
	// fmt.Println(board)
	fmt.Println()
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)
}
