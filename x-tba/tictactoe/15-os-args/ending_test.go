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
