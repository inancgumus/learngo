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

const maxTurns = 9

var (
	won    bool
	turn   int
	player string // current player
)

func main() {
	player = player1
	turn = maxTurns

	// you need to call the other functions but the main
	printBoard()
	printStatus()
	printEnding()
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

	switch {
	case won:
		fmt.Printf(">>> WINNER: %s\n", player)
	case turn == maxTurns:
		fmt.Println(">>> TIE!")
	default:
		switchPlayer()
		fmt.Println(">> NEXT TURN!", player)
	}
}

// switchPlayer switches to the next player
func switchPlayer() {
	// switch the player
	if player == player1 {
		player = player2
	} else {
		player = player1
	}
}
