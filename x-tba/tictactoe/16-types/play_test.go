package main

import "testing"

func TestWrongMove(t *testing.T) {
	// /---+---+---\
	// | X | X | X |
	// +---+---+---+
	// | X | X | X |
	// +---+---+---+
	// | X | X | X |
	// \---+---+---/

	// fill the board with artificial cells
	for i := range cells {
		cells[i] = player1
	}

	// any move beyond this point is wrong.
	// reason: every cell is occupied.
	if play(); !wrongMove {
		t.Errorf("wrongMove = %t; want true", wrongMove)
	}

	initCells()
}
