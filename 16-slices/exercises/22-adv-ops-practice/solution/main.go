// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10
	s.Width = 60

	//
	// #1
	//

	names := make([]string, 5)
	s.Show("1st step", names)

	//
	// #2
	//

	names = append(names, "einstein", "tesla", "aristo")
	s.Show("2nd step", names)

	//
	// #3
	//
	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristo")
	s.Show("3rd step", names)

	//
	// #4
	//

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]

	s.Show("4th step", names)

	//
	// #5
	//

	clone := make([]string, 3, 5)
	copy(clone, names[len(names)-3:])
	s.Show("5th step (before append)", clone)

	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	//
	// #6
	//

	sliced := clone[1:4:4]
	sliced = append(sliced, "hypatia")

	clone[2] = "elder"

	s.Show("6th step", clone, sliced)
}
