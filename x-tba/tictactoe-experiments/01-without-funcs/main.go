// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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

	board = [3][3]string{
		{empty, empty, empty},
		{empty, empty, empty},
		{empty, empty, empty},
	}

	player = player1

	// for saving the replay.log
	inputs []byte
)

func main() {
	rainbow.Rainbow(banner, strings.Count(banner, "\n"))

	in := bufio.NewScanner(os.Stdin)
	for {
		// -------------------------------------------------
		// PRINT THE BOARD AND THE PROMPT
		// -------------------------------------------------
		fmt.Printf("\n %s\n", header)
		for _, line := range board {
			fmt.Printf(" %s\n", strings.Join(line[:], separator))
			fmt.Printf(" %s\n", footer)
		}

		// -------------------------------------------------
		// IS THERE A WINNER? OR IS IT A TIE?
		// -------------------------------------------------
		for _, m := range [2]string{player1, player2} {
			b, mmm := board, strings.Repeat(m, 3)

			won = /* horizontals */
				strings.Join(b[0][:], "") == mmm ||
					strings.Join(b[1][:], "") == mmm ||
					strings.Join(b[2][:], "") == mmm ||

					/* verticals */
					b[0][0]+b[1][0]+b[2][0] == mmm ||
					b[0][1]+b[1][1]+b[2][1] == mmm ||
					b[0][2]+b[1][2]+b[2][2] == mmm ||

					/* diagonals */
					b[0][0]+b[1][1]+b[2][2] == mmm ||
					b[0][2]+b[1][1]+b[2][0] == mmm

			if won {
				player = m
				break
			}
		}

		if won {
			player = strings.TrimSpace(player)
			fmt.Printf("\n>>> WINNER: %s\n", player)
			break
		} else if turn++; turn == maxTurn+1 {
			fmt.Printf("\n>>> TIE!\n")
			break
		}

		// -------------------------------------------------
		// CHECK THE MOVE AND PLAY
		// -------------------------------------------------

		// get the input
		fmt.Printf("\nPLAYER: %q [1-9]: ", player)
		if !in.Scan() {
			break
		}

		// Atoi already return 0 on error; no need to check
		// it for the following switch to work
		pos, _ := strconv.Atoi(in.Text())

		// Save the input for replaying
		inputs = append(inputs, byte(pos+48), '\n')

		var row int
		switch {
		case pos >= 1 && pos <= 3:
			row = 0
		case pos >= 4 && pos <= 6:
			row = 1
		case pos >= 7 && pos <= 9:
			row = 2
		default:
			fmt.Println("\n>>>", "wrong position!")
			continue
		}

		col := pos - row*3 - 1

		if board[row][col] != empty {
			fmt.Println("\n>>>", "already played!")
			continue
		}

		// put a mark on the board
		board[row][col] = player

		// switch to the next player
		if player == player1 {
			player = player2
		} else {
			player = player1
		}
	}

	if err := ioutil.WriteFile("replay.log", inputs, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot save replay: %v", err)
	}
}
