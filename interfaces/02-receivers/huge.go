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
)

type huge struct {
	// about ~200mb
	games [10_000_000]game

	// if this syntax doesn't work on your system, type it as:
	// games [10000000]game
}

// only copies a single pointer.
func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}

// each time it is called,
// the original value will be copied.
// calling addr() x 10 times = ~2 GB.
// func (h huge) addr() {
// 	fmt.Printf("%p\n", &h)
// }
