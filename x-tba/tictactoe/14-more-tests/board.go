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
