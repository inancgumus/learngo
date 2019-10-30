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

func main() {
	fmt.Println(`
~~~~~~~~~~~~~~~
  TIC~TAC~TOE 
~~~~~~~~~~~~~~~`)

	loop(bufio.NewScanner(os.Stdin))
}

const (
	// skin options :-)
	emptyMark = "☆"
	mark1     = "💀"
	mark2     = "🎈"

	maxTurn = 9
)

type game struct {
	board  [][]string
	turn   int
	player string

	buf strings.Builder
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

		if msg := g.finito(); msg != "" {
			fmt.Printf("\n%s%s\n", g, msg)
			break
		}

		g.next()
	}
}

func newGame() *game {
	return &game{
		player: mark1,
		board: [][]string{
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
			{emptyMark, emptyMark, emptyMark},
		},
	}
}

func (g *game) prompt() {
	fmt.Printf("\n%s%s [1-9]: ", g, g.player)
}

func (g *game) String() string {
	g.buf.Reset()

	g.buf.WriteRune('\n')
	g.buf.WriteString("---+----+---\n")

	for _, line := range g.board {
		g.buf.WriteString(strings.Join(line, "  | "))
		g.buf.WriteRune('\n')

		g.buf.WriteString("---+----+---\n")
	}

	return g.buf.String()
}

func (g *game) play(pos int) string {
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

func (g *game) finito() string {
	g.turn++

	switch {
	case g.won():
		return fmt.Sprintf("WINNER: %s", g.player)
	case g.turn == maxTurn:
		return "TIE!"
	}
	return ""
}

func (g *game) won() (won bool) {
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

func (g *game) next() {
	if g.player == mark1 {
		g.player = mark2
	} else {
		g.player = mark1
	}
}
