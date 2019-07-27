package main

// play plays the game for the current player.
// + registers the player's move in the board.
// if the move is valid:
// + increases the turn.
func play() {
	// [" ", " ", " ", " ", " ", " ", " ", " ", " " ] -> cells slice
	//   0    1    2    3    4    5    6    7    8    -> indexes

	// /---+---+---\
	// | 0 | 1 | 2 |
	// +---+---+---+
	// | 3 | 4 | 5 |
	// +---+---+---+
	// | 6 | 7 | 8 |
	// \---+---+---/

	// current player plays to the 4th position (index)
	// play to a fixed position "for now".
	pos := 4

	// register the move: put the player's sign on the board
	cells[pos] = player

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
