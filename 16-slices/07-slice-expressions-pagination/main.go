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

	// s.Show("0:4", items[0:4])
	// s.Show("4:8", items[4:8])
	// s.Show("8:12", items[8:12])
	// s.Show("12:13", items[12:13])
	// s.Show("12:14", items[12:14]) // error

	l := len(items)

	const pageSize = 4

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}

		// fmt.Printf("%d:%d\n", from, to)

		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		s.Show(head, currentPage)
	}
}
