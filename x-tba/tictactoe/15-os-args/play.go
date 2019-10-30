// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "math/rand"

// play plays the game for the current player.
// + registers the player's move in the board.
// if the move is valid:
// + increases the turn.
func play() {
	// pick a random move (very intelligent AI!)
	// it can play to the same position!
	lastPos = rand.Intn(maxTurns)

	// is it a valid move?
	if cells[lastPos] != emptyCell {
		wrongMove = true

		// skip the rest of the function from running
		return
	}

	// register the move: put the player's sign on the board
	cells[lastPos] = player

	// increment the current turns
	turn++
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
