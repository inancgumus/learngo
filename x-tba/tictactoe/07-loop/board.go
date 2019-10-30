// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// printBoard prints the board
func printBoard() {
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

	const cellsPerLine = 3

	// Prevents the Println '\n' warning
	fmt.Printf("%s\n\n", banner)
	fmt.Printf("%s\n", sepHeader)

	for i := 0; i < maxTurns; i++ {
		lineStarts := i%cellsPerLine == 0
		lineEnds := (i+1)%cellsPerLine == 0
		lastLine := (i + 1) == maxTurns

		if lineStarts {
			fmt.Print(sepCell)
		}

		switchPlayer()
		fmt.Printf(" %s ", player)

		if lineEnds {
			fmt.Println(sepCell)
		}

		if lineEnds && !lastLine {
			fmt.Println(sepLine)
		}

		if lineEnds {
			continue // prevents printing the sepCell below
		}

		// print the cell separator until 3 cells per line
		fmt.Print(sepCell)
	}

	fmt.Printf("%s\n", sepFooter)
}
