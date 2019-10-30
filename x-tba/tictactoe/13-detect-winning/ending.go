// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// -------------------------------------------------
// IS THERE A WINNER? OR IS IT A TIE?
// -------------------------------------------------

// /---+---+---\
// | 0 | 1 | 2 |
// +---+---+---+
// | 3 | 4 | 5 |
// +---+---+---+
// | 6 | 7 | 8 |
// \---+---+---/

func checkWinOrTie() {
	// intentional bug: tie shouldn't happen before winning detection
	if tie = turn == maxTurns; tie {
		// if tie don't check for the winning
		return
	}

	// loop over all the players
	for i := 1; i <= 2; i++ {
		// check for the next player
		p := player2
		if i == 1 {
			p = player1
		}

		/* check horizontals */
		hor := (cells[0] == p && cells[1] == p && cells[2] == p) ||
			(cells[3] == p && cells[4] == p && cells[5] == p) ||
			(cells[6] == p && cells[7] == p && cells[8] == p)

		/* check verticals */
		ver := (cells[0] == p && cells[3] == p && cells[6] == p) ||
			(cells[1] == p && cells[4] == p && cells[7] == p) ||
			(cells[2] == p && cells[5] == p && cells[8] == p)

		/* check diagonals */
		diag := (cells[0] == p && cells[4] == p && cells[8] == p) ||
			(cells[2] == p && cells[4] == p && cells[6] == p)

		// any winner?
		if hor || ver || diag {
			won = true

			// this player wins
			player = p

			// there is a winner so don't check for tie!
			return
		}
	}
}
