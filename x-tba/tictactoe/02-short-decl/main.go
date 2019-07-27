package main

import "fmt"

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/
func main() {
	var won bool // is there any winner?

	banner := " TIC~TAC~TOE"
	board := `
/---+---+---\
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
\---+---+---/`

	// short declaration (type-inference)
	maxTurns, turn := 9, 0
	progress := (1 - (float64(turn) / float64(maxTurns))) * 100

	fmt.Println(banner)
	fmt.Println(board)
	fmt.Println()
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)
}
