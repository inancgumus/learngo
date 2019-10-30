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

type game struct {
	board  [][]string
	turn   int
	player string
}

func loop(in *bufio.Scanner) {
	g := newGame()

	for {
		g.prompt()

		if !in.Scan() {
			break
		}

		// Atoi already return 0 on error; no need to check
		// it for the following switch to work
		pos, _ := strconv.Atoi(in.Text())

		msg := g.play(pos)
		if msg != "" {
			fmt.Println("\n>>>", msg)
			continue
		}

		g.turn++
		if msg := g.finito(); msg != "" {
			g.print()
			fmt.Printf("\n%s\n", msg)
			break
		}

		g.player = switchTo(g.player)
	}
}

func newGame() game {
	return game{
		player: mark1,
		board: [][]string{
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
		},
	}
}

func (g game) prompt() {
	g.print()
	fmt.Printf("\n%s [1-9]: ", g.player)
}

func (g game) print() {
	fmt.Println()
	fmt.Println("---+----+---")
	for _, line := range g.board {
		fmt.Printf("%s\n", strings.Join(line, "  | "))
		fmt.Println("---+----+---")
	}
}

func (g game) play(pos int) string {
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

	if g.board[row][col] != emptyMark {
		return "already played!"
	}

	// put the player
	g.board[row][col] = g.player

	return ""
}

func (g game) finito() string {
	switch {
	case g.won():
		return fmt.Sprintf("WINNER: %s", g.player)
	case g.turn == maxTurn:
		return "TIE!"
	}
	return ""
}

func (g game) won() (won bool) {
	for _, m := range [2]string{mark1, mark2} {
		b, mmm := g.board, strings.Repeat(m, 3)

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
