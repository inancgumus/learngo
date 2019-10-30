// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// Examples are normally used for showing how to use your package.
// But you can also use them as output testing.

func ExamplePrintBoard() {
	// let the printBoard function print an output

	printBoard()

	// the output should exactly match the following (after Output:)

	// Output:
	// ~~~~~~~~~~~~~~~
	//   TIC~TAC~TOE
	// ~~~~~~~~~~~~~~~
	//
	// /---+---+---\
	// | X | O | X |
	// +---+---+---+
	// | O | X | O |
	// +---+---+---+
	// | X | O | X |
	// \---+---+---/
}
