// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	rainbow "github.com/guineveresaenger/golang-rainbow"
)

const (
	maxTurn = 9

	// skin options :-)
	empty     = "   "
	player1   = " X "
	player2   = " O "
	header    = "---+---+---"
	footer    = "---+---+---"
	separator = "|"

	banner = `
~~~~~~~~~~~~~~~
  TIC~TAC~TOE 
~~~~~~~~~~~~~~~`
)

// -------------------------------------------------
// INITIALIZE THE GAME
// -------------------------------------------------
var (
	turn int
	won  bool

	board = [][]string{
		{empty, empty, empty},
		{empty, empty, empty},
		{empty, empty, empty},
	}

	player = player1
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rainbow.Rainbow(banner, strings.Count(banner, "\n"))

	for {
		// -------------------------------------------------
		// PRINT THE BOARD AND THE PROMPT
		// -------------------------------------------------
		fmt.Printf("\n %s\n", header)
		for _, line := range board {
			fmt.Printf(" %s\n", strings.Join(line, separator))
			fmt.Printf(" %s\n", footer)
		}

		// -------------------------------------------------
		// IS THERE A WINNER? OR IS IT A TIE?
		// -------------------------------------------------
		for i := 1; i <= 2; i++ {
			m := player2
			if i == 1 {
				m = player1
			}

			b, mmm := board, strings.Repeat(m, 3)

			/* horizontals */
			hor := strings.Join(b[0], "") == mmm ||
				strings.Join(b[1], "") == mmm ||
				strings.Join(b[2], "") == mmm

				/* verticals */
			ver := b[0][0]+b[1][0]+b[2][0] == mmm ||
				b[0][1]+b[1][1]+b[2][1] == mmm ||
				b[0][2]+b[1][2]+b[2][2] == mmm

				/* diagonals */
			diag := b[0][0]+b[1][1]+b[2][2] == mmm ||
				b[0][2]+b[1][1]+b[2][0] == mmm

			won = hor || ver || diag

			if won {
				player = m
				break
			}
		}

		if won {
			player = strings.TrimSpace(player)
			fmt.Printf("\n>>> WINNER: %s\n", player)
			break
		} else if turn == maxTurn {
			fmt.Printf("\n>>> TIE!\n")
			break
		}

		// -------------------------------------------------
		// PLAY
		// -------------------------------------------------

		pos := rand.Intn(9) + 1
		fmt.Printf("\nPlayer %s plays to %d\n", player, pos)

		// -------------------------------------------------
		// IS IT A VALID MOVE?
		// -------------------------------------------------

		var row int
		switch {
		case pos <= 3:
			row = 0
		case pos <= 6:
			row = 1
		case pos <= 9:
			row = 2
		default:
			fmt.Println(">>>", "wrong position!")
			continue
		}

		col := pos - row*3 - 1

		if board[row][col] != empty {
			fmt.Println(">>>", "already played!")
			continue
		}

		// -------------------------------------------------
		// MARK THE MOVE AND INCREMENT THE TURN
		// -------------------------------------------------

		// put a mark on the board
		board[row][col] = player

		// switch to the next player
		if player == player1 {
			player = player2
		} else {
			player = player1
		}

		turn++
		// time.Sleep(time.Millisecond * 100)
	}
}
