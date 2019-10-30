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
)

func main() {
	fmt.Println(banner)

	loop(bufio.NewScanner(os.Stdin))
}

func loop(in *bufio.Scanner) {
	var (
		turn   int
		player = mark1
		board  = createBoard()
	)

	for {
		prompt(board, player)

		if !in.Scan() {
			break
		}

		msg := play(board, player, getMove(in.Text()))
		if msg != "" {
			fmt.Println("\n>>>", msg)
			continue
		}

		turn++
		if msg := finito(board, turn, player); msg != "" {
			printBoard(board, player)
			fmt.Printf("\n%s\n", msg)
			break
		}

		player = switchTo(player)
	}
}

func createBoard() [][]string {
	return [][]string{
		{emptyMark, emptyMark, emptyMark},
		{emptyMark, emptyMark, emptyMark},
		{emptyMark, emptyMark, emptyMark},
	}
}

func prompt(board [][]string, player string) {
	printBoard(board, player)
	fmt.Printf("\n%s [1-9]: ", player)
}

func printBoard(board [][]string, player string) {
	fmt.Println()
	fmt.Println("---+----+---")
	for _, line := range board {
		fmt.Printf("%s\n", strings.Join(line, "  | "))
		fmt.Println("---+----+---")
	}
}

func getMove(move string) int {
	// Atoi already return 0 on error; no need to check
	// it for the following switch to work
	pos, _ := strconv.Atoi(move)
	return pos
}

func play(board [][]string, player string, pos int) string {
	var row int

	switch {
	case pos >= 7 && pos <= 9:
		row = 2
	case pos >= 4 && pos <= 6:
		row = 1
	case pos >= 1 && pos <= 3:
		row = 0
	default:
		return "wrong position"
	}

	col := pos - row*3 - 1

	if board[row][col] != emptyMark {
		return "already played!"
	}

	// put the player
	board[row][col] = player

	return ""
}

func finito(board [][]string, turn int, player string) string {
	switch {
	case won(board):
		return fmt.Sprintf("WINNER: %s", player)
	case turn == maxTurn:
		return "TIE!"
	}
	return ""
}

func won(board [][]string) (won bool) {
	for _, m := range [2]string{mark1, mark2} {
		b, mmm := board, strings.Repeat(m, 3)

		won = /* horizontals */
			strings.Join(b[0], "") == mmm ||
				strings.Join(b[1], "") == mmm ||
				strings.Join(b[2], "") == mmm ||

				/* verticals */
				b[0][0]+b[1][0]+b[2][0] == mmm ||
				b[0][1]+b[1][1]+b[2][1] == mmm ||
				b[0][2]+b[1][2]+b[2][2] == mmm ||

				/* diagonals */
				b[0][0]+b[1][1]+b[2][2] == mmm ||
				b[0][2]+b[1][1]+b[2][0] == mmm

		if won {
			return true
		}
	}
	return false
}

func switchTo(player string) string {
	if player == mark1 {
		return mark2
	}
	return mark1
}
