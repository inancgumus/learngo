// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// think of this as search results of a search engine.
	// it could have been fetched from a database
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("All items", items)

	top3 := items[:3]
	s.Show("Top 3 items", top3)

	l := len(items)

	// you can use variables in a slice expression
	last4 := items[l-4:]
	s.Show("Last 4 items", last4)

	// reslicing: slicing another sliced slice
	mid := last4[1:3]
	s.Show("Last4[1:3]", mid)

	// the same elements can be in different indexes
	// fmt.Println(items[9], last4[0])

	// slicing returns a slice with the same type of the sliced slice
	fmt.Printf("slicing : %T %[1]q\n", items[2:3])

	// indexing returns a single element with the type of the indexed slice's element type
	fmt.Printf("indexing: %T %[1]q\n", items[2])
}
