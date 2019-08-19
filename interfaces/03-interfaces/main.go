// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)

	tetris.discount(.8)
	// printer(tetris).discount(.8)
	store.print()

	// type printer interface {
	// 	print()
	// }

	// p can store a value of any type that has a `print()` method
	var p printer
	p = puzzle{title: "sidewinder", price: 10}
	p.print()
}