// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// cell represents a tictactoe board cell
type cell string

// skin options :-)
const (
	banner = `
~~~~~~~~~~~~~~~
  TIC~TAC~TOE
~~~~~~~~~~~~~~~`

	player1, player2, emptyCell cell = "X", "O", " "

	sepHeader = `/---+---+---\`
	sepLine   = `+---+---+---+`
	sepFooter = `\---+---+---/`
	sepCell   = "|"
)
