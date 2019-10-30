// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "testing"

func TestWin(t *testing.T) {
	// /---+---+---\
	// | O | X | X |
	// +---+---+---+
	// | O | O | X |
	// +---+---+---+
	// | X | O | X |
	// \---+---+---/

	cells[0] = player2
	cells[1] = player1
	cells[2] = player1
	cells[3] = player2
	cells[4] = player2
	cells[5] = player1
	cells[6] = player1
	cells[7] = player2
	cells[8] = player1
	turn = maxTurns

	if checkWinOrTie(); !won {
		t.Errorf("won = %t; want true", won)
	}

	// TestWin was messing up with the results.
	// Test ordering shouldn't be important.
	initCells()
}

func TestTie(t *testing.T) {
	// /---+---+---\
	// | O | X | X |
	// +---+---+---+
	// | X | O | O |
	// +---+---+---+
	// | X | O | X |
	// \---+---+---/

	cells[0] = player2
	cells[1] = player1
	cells[2] = player1
	cells[3] = player1
	cells[4] = player2
	cells[5] = player2
	cells[6] = player1
	cells[7] = player2
	cells[8] = player1
	turn = maxTurns

	if checkWinOrTie(); !tie {
		t.Errorf("tie = %t; want true", tie)
	}

	initCells()
}
