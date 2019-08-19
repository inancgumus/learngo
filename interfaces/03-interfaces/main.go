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
}

/*
var my list

// + only the `*game` have methods.
// + `game` doesn't have any methods.
// + `item` interface couldn't be satisfied with a `game` value.
// + it can only be satisfied with a `*game` value.
// my = my.add(minecraft)
my = append(my, &minecraft)
my = append(my, &tetris)
my = append(my, &mobydick)

// interface value stores a pointer to minecraft
minecraft.discount(.5)
my.list()

my.discount(.5)
my.list()
*/

// ! cannot add mobydick to the store: it's a book type, not a game type.
// interfaces to the rescue!
// mobydick := book{title: "moby dick", price: 10}
// my.add(&mobydick)
// ----------------------------------------------------------------
