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
		// mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)

	// you can attach methods to a compatible type on the fly:
	// items -> []*game
	// list  -> []*game
	my := list(items)
	// my = nil

	// you can call methods even on a nil value
	my.print()
}
