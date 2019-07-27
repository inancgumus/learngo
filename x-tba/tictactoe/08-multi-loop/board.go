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

	fmt.Printf("%s\n\n", banner)
	fmt.Printf("%s\n", sepHeader)

	for i, step := 0, 3; i < maxTurns; i += step {
		fmt.Print(sepCell)

		for j := 0; j < step; j++ {
			switchPlayer()
			fmt.Printf(" %s %s", player, sepCell)
		}
		fmt.Println()

		if lastLine := (i + step); lastLine != maxTurns {
			fmt.Println(sepLine)
		}
	}

	fmt.Printf("%s\n", sepFooter)
}
