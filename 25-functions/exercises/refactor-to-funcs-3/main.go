// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #3
//
//  Let's continue from the previous exercise. This time,
//  you're going to add json encoding and decoding using
//  funcs.
//
//  1- Create a new command in a func that encodes the
//     game store data to json and terminates the program.
//     Lastly, add it to runCmd func.
//
//     For more information, see: "Encode" exercise from
//     the structs section.
//
//        Name  : cmdSave
//        Inputs: []game
//        Output: bool
//
//  2- Refactor the load() to load the games from the
//     `data` constant (it's in the games.go as well).
//
//     For more information, see: "Decode" exercise from
//     the structs section.
//
// ---------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	games := load()
	byID := indexByID(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits

`)

		if !in.Scan() || !runCmd(in.Text(), games, byID) {
			break
		}
	}
}
