// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
		yoda      = toy{title: "yoda", price: 150}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik, &yoda)

	// #2
	store.discount(.5)
	store.print()

	// #1
	// var p printer
	// p = &tetris
	// tetris.discount(.5)
	// p.print()
}
