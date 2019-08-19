// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

func main() {
	var (
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		mobydick  = book{title: "moby dick", price: 10}
	)

	var my list
	// my.list() // nil receiver is a valid value

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
}
