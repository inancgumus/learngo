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
	fmt.Printf("%s\n\n", banner)
	fmt.Printf("%s\n", sepHeader)

	for i, step := 0, 3; i < len(cells); i += step {
		fmt.Print(sepCell)

		for j := 0; j < step; j++ {
			move := cells[i+j]
			fmt.Printf(" %s %s", move, sepCell)
		}

		fmt.Println()

		if lastLine := (i + step); lastLine != len(cells) {
			fmt.Println(sepLine)
		}
	}

	fmt.Printf("%s\n", sepFooter)
}
