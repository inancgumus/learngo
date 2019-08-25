// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// import "fmt"

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10, readTime: 10}
		minecraft = game{title: "minecraft", price: 20, playTime: 5}
		tetris    = game{title: "tetris", price: 5, playTime: 2}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)
	// tetris.discount(.8)
	// store.print()

	store.discount(.5)
	store.print()

	// t := store.time()
	// fmt.Printf("Total entertainment time: %d hours\n", t)

	// games := store[:2]
	// other := store[2:]
	// games.print()
	// other.print()
	// list{games, other}.print()

	// var b *book

	// var np printer
	// np = b
	// // p := printer(puzzle{title: "sidewinder", price: 10})
	// fmt.Println(np == nil)
}
