// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #4: method expressions
	// methods are just functions that receive a value of a type.
	game.print(minecraft) // sends `minecraft` value to `game.print`
	game.print(tetris)    // sends `tetris` value to `game.print`
	fmt.Println()

	// #3
	mobydick.print()  // sends `mobydick` value to `book.print`
	minecraft.print() // sends `minecraft` value to `game.print`
	tetris.print()    // sends `tetris` value to `game.print`

	// #2
	// mobydick.printBook()
	// minecraft.printGame()

	// #1
	// printBook(mobydick)
	// printGame(minecraft)
}
