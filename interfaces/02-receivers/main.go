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
	)

	// sends a pointer of minecraft to the discount method
	// same as: (&minecraft).discount(.1)
	minecraft.discount(.1)

	mobydick.print()
	minecraft.print()
	tetris.print()

	// creates a variable that occupies 200mb on memory
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
