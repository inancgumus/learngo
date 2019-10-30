// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// nextTurn prints the board for the next turn and checks for the winning conditions.
// if win or tie: returns false, otherwise true.
func nextTurn() bool {
	play()
	printBoard()

	fmt.Printf("\n>>> PLAYER %q PLAYS to %d\n", player, lastPos+1)

	// the switch below is about winning and tie conditions.
	// so it is good have checkWinOrTie() as a simple statement.
	// totally optional.
	switch checkWinOrTie(); {
	default:
		switchPlayer()
		printStatus()

	case wrongMove:
		fmt.Printf(">>> CELL IS OCCUPIED: PLAY AGAIN!\n")
		wrongMove = false // reset for the next turn

	case won, tie:
		if won {
			fmt.Println(">>> WINNER:", player)
		} else {
			fmt.Println(">>> TIE!")
		}
		return false
	}
	return true
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
