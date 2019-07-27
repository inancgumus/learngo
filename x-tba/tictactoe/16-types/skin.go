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
