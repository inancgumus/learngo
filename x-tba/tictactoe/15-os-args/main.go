package main

import (
	"fmt"
	"time"
)

/*
~ TICTACTOE GAME IN GO ~
+ This example uses the very basics of the Go language.
+ The goal is learning all the basics.
*/

const (
	maxTurns         = 9
	defaultGameSpeed = time.Second * 2 // time between plays
)

var (
	won, tie bool // is there any winner or a tie?
	turn     int  // total valid turns played

	cells     [maxTurns]string // used to draw the board: contains the players' moves
	lastPos   int              // last played position
	wrongMove bool             // was the last move wrong?

	player    = player1          // current player
	gameSpeed = defaultGameSpeed // sets the default game speed
)

// main is only responsible for the game loop, that's it.
func main() {
	setGameSpeed()
	printBoard()
	wait()

	for nextTurn() {
		wait()
	}
}

func wait() {
	fmt.Println()
	time.Sleep(gameSpeed) // player thinks...
}
