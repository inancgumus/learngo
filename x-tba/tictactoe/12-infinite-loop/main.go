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
	won    bool      // is there any winner?
	turn   int       // total valid turns played
	player = player1 // current player

	cells     [maxTurns]string // used to draw the board: contains the players' moves
	lastPos   int              // last played position
	wrongMove bool             // was the last move wrong?
)

func main() {
	printBoard()

	// alternative:
	// for won || tie {}

	// this a label: it can be used by break, continue and goto
theGameLoop:
	// loop forever until the game ends (tie or win)
	for {
		wait()
		play()
		printBoard()

		fmt.Printf("\n>>> PLAYER %q PLAYS to %d\n", player, lastPos+1)

		// simple statement
		switch tie := turn == maxTurns; {
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
			break theGameLoop
		}
	}
}

func wait() {
	fmt.Println()
	fmt.Scanln()
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
