package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// skin options :-)
	emptyMark = "☆"
	mark1     = "💀"
	mark2     = "🎈"
	banner    = `
~~~~~~~~~~~~~~~
  TIC~TAC~TOE 
~~~~~~~~~~~~~~~`

	maxTurn = 9

	prompt = "\n%s [1-9]: "
)

func main() {
	// -------------------------------------------------
	// INITIALIZE THE GAME
	// -------------------------------------------------

	fmt.Println(banner)

	in := bufio.NewScanner(os.Stdin)

	var (
		turn   int
		player = mark1
		board  = [3][3]string{
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
		}
	)

game:
	for {
		// -------------------------------------------------
		// PRINT THE BOARD AND THE PROMPT
		// -------------------------------------------------
		fmt.Println()
		fmt.Println("---+----+---")
		for _, line := range board {
			fmt.Printf("%s\n", strings.Join(line[:], "  | "))
			fmt.Println("---+----+---")
		}
		fmt.Printf(prompt, player)

		// get the input
		if !in.Scan() {
			break
		}

		// -------------------------------------------------
		// CHECK THE MOVE AND PLAY
		// -------------------------------------------------
		var (
			pos      int
			row, col int
		)

		// Atoi already return 0 on error; no need to check
		// it for the following switch to work
		pos, _ = strconv.Atoi(in.Text())

		switch {
		case pos >= 7 && pos <= 9:
			row = 2
		case pos >= 4 && pos <= 6:
			row = 1
		case pos >= 1 && pos <= 3:
			row = 0
		default:
			fmt.Println("\n>>>", "wrong position!")
			continue
		}

		col = pos - row*3 - 1

		if board[row][col] != emptyMark {
			fmt.Println("\n>>>", "already played!")
			continue
		}

		// put a mark on the board
		board[row][col] = player

		// -------------------------------------------------
		// IS THERE A WINNER? OR IS IT A TIE?
		// -------------------------------------------------
		var won bool

		for _, m := range [2]string{mark1, mark2} {
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
				break
			}
		}

		if won {
			fmt.Printf("\nWINNER: %s\n", player)
			break game
		} else if turn++; turn == maxTurn {
			fmt.Printf("\nTIE!\n")
			break game
		}

		// -------------------------------------------------
		// CHANGE THE MARKER TO THE NEXT PLAYER
		// -------------------------------------------------
		if player == mark1 {
			player = mark2
		} else {
			player = mark1
		}
	}
}
