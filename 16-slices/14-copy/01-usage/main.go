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
	evens := []int{2, 4}
	odds := []int{3, 5, 7}

	s.Show("evens [before]", evens)
	s.Show("odds  [before]", odds)

	N := copy(evens, odds)
	fmt.Printf("%d element(s) are copied.\n", N)

	s.Show("evens [after]", evens)
	s.Show("odds  [after]", odds)
}
