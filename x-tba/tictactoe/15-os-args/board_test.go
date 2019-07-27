package main

// Examples are normally used for showing how to use your package.
// But you can also use them as output testing.

func ExamplePrintBoard() {
	printBoard()

	// Output:
	// ~~~~~~~~~~~~~~~
	//   TIC~TAC~TOE
	// ~~~~~~~~~~~~~~~
	//
	// /---+---+---\
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// \---+---+---/
}

func ExamplePrintBoardCells() {
	cells[0] = player1
	cells[4] = player2
	cells[8] = player1
	printBoard()

	// Output:
	// ~~~~~~~~~~~~~~~
	//   TIC~TAC~TOE
	// ~~~~~~~~~~~~~~~
	//
	// /---+---+---\
	// | X |   |   |
	// +---+---+---+
	// |   | O |   |
	// +---+---+---+
	// |   |   | X |
	// \---+---+---/
}
