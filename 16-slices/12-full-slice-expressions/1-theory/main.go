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

	sliceable := []byte{'f', 'u', 'l', 'l'}

	s.Show("sliceable", sliceable)
	s.Show("sliceable[0:3]", sliceable[0:3])
	s.Show("sliceable[0:3:3]", sliceable[0:3:3])
	s.Show("sliceable[0:2:2]", sliceable[0:2:2])
	s.Show("sliceable[0:1:1]", sliceable[0:1:1])
	s.Show("sliceable[1:3:3]", sliceable[1:3:3])
	s.Show("sliceable[2:3:3]", sliceable[2:3:3])
	s.Show("sliceable[2:3:4]", sliceable[2:3:4])
	s.Show("sliceable[4:4:4]", sliceable[4:4:4])
}
