// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	// ----------------------------------------------------
	// #1 nil slice
	var games []string // nil slice
	s.Show("games", games)

	// ----------------------------------------------------
	// #2 empty slice
	games = []string{} // empty slice
	s.Show("games", games)
	// s.Show("another empty", []int{})

	// ----------------------------------------------------
	// #3 non-empty slice
	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)

	// ----------------------------------------------------
	// #4 reset the part using the games slice
	//    part is empty but its cap is still 4
	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part[:0]", part)
	s.Show("part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	// #6 backing array's elements become inaccessible
	// games = games[len(games):]

	// ----------------------------------------------------
	// #5 part doesn't have any more capacity
	//    games slice is still intact
	s.Show("part", part)
	s.Show("games", games)
}
