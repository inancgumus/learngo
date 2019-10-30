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

const (
	banner = `
 TIC~TAC~TOE

/---+---+---\
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
\---+---+---/`

	maxTurns = 9
)

var (
	won  bool
	turn int
)

func main() {
	// you need to call the other functions but the main
	printBoard()
	printStatus()
	printEnding()
}

// printBoard cannot use the banner
func printBoard() {
	// it can only use the package-level names (identifiers)
	fmt.Println(banner)
}

// printStatus prints the current status of the game
// it cannot access to the names (vars, consts, etc) inside any other func
func printStatus() {
	fmt.Println()

	progress := (1 - (float64(turn) / maxTurns)) * 100
	fmt.Printf("Current Turn           : %d\n", turn)
	fmt.Printf("Is there a winner      : %t\n", won)
	fmt.Printf("Turns left             : %.1f%%\n", progress)
}

func printEnding() {
	fmt.Println()

	// creates a new variable that belongs to printEnding()
	// current player
	player := "X"

	// accesses to the package-level vars: won, turn
	// won = true
	// turn = maxTurns

	// if won {
	// 	fmt.Printf(">>> WINNER: %s\n", player)
	// } else if turn == maxTurns { // != >= <= < >
	// 	fmt.Println(">>> TIE!")
	// } else {
	// 	fmt.Println(">> NEXT TURN!")
	// }

	switch {
	case won:
		// player := "X"
		fmt.Printf(">>> WINNER: %s\n", player)
	case turn == maxTurns:
		fmt.Println(">>> TIE!")
	default:
		fmt.Println(">> NEXT TURN!", player)

	}

	// invalid if the player is within the scope of the switch:
	// fmt.Println(">> Current Player: ", player)
}
