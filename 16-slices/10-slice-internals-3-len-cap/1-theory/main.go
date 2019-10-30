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
	s.MaxPerLine = 6
	s.PrintBacking = true

	ages := []int{35, 15, 25}
	s.Show("ages", ages)

	s.Show("ages[0:0]", ages[0:0])

	for i := 1; i < 4; i++ {
		txt := fmt.Sprintf("ages[%d:%d]", 0, i)
		s.Show(txt, ages[0:i])
	}

	s.Show("append", append(ages, 50))
}
