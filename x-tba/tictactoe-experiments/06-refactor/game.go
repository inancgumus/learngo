// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"strings"
)

type state int

const (
	maxTurn = 9

	wrongPosition = -2

	statePlaying state = iota
	stateWon
	stateTie

	stateAlreadyPlayed
	stateWrongPosition
)

type game struct {
	board [][]string

	turn   int
	player string

	skin   // embed the skin
	logger // embed the logger
}

func newGame(s skin, l logger) *game {
	return &game{
		player: s.mark1,
		board: [][]string{
			{s.empty, s.empty, s.empty},
			{s.empty, s.empty, s.empty},
			{s.empty, s.empty, s.empty},
		},
		skin:   s,
		logger: l,
	}
}

func (g *game) play(pos int) state {
	if st := g.move(pos); st != statePlaying {
		return st
	}

	g.turn++ // increment the turn

	// first check the winner then check the tie
	// or the last mover won't win
	switch {
	case g.won():
		return stateWon
	case g.turn == maxTurn:
		return stateTie
	}

	g.changePlayer()
	return statePlaying
}

func (g *game) move(pos int) state {
	row, col := position(pos)
	if row+col == wrongPosition {
		return stateWrongPosition
	}

	if g.board[row][col] != g.empty {
		return stateAlreadyPlayed
	}

	// put the player's mark on the board
	g.board[row][col] = g.player

	return statePlaying
}

// we can detect the winning state just by comparing the strings
// because, the game board is a bunch of strings
func (g *game) won() (won bool) {
	for _, m := range [2]string{g.mark1, g.mark2} {
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

// this method should have a pointer receiver
// because, it changes the game value
func (g *game) changePlayer() {
	if g.player == g.mark1 {
		g.player = g.mark2
	} else {
		g.player = g.mark1
	}
}

func (g *game) print() {
	g.Println()
	g.Println(g.header)

	for i, line := range g.board {
		g.Print(g.separator)

		for _, m := range line {
			g.Printf("%2s%s", m, g.separator)
		}

		if i+1 != len(g.board) {
			g.Printf("\n%s\n", g.middle)
		}
	}

	g.Printf("\n%s\n", g.footer)
}

// this function doesn't depend on the game state
// so, make it a function instead of a method
func position(pos int) (row, col int) {
	switch {
	case pos >= 1 && pos <= 3:
		row = 0
	case pos >= 4 && pos <= 6:
		row = 1
	case pos >= 7 && pos <= 9:
		row = 2
	default:
		return -1, -1
	}

	return row, pos - row*3 - 1
}
