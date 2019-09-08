// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	var store list

	db := database{list: &store}

	// db.register("book", new(book))
	// db.register("book", new(book))
	// db.register("game", new(game))
	// db.register("puzzle", new(puzzle))
	// db.register("toy", new(toy))

	// db.load("database.json")

	store = list{
		&book{product{"moby dick", 10}, toTimestamp(118281600)},
		&book{product{"odyssey", 15}, toTimestamp("733622400")},
		&book{product{"hobbit", 25}, unknown},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	db.save("database.json")

	fmt.Print(store)
}
