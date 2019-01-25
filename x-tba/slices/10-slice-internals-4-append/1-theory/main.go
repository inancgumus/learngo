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

	ages := []int{35, 15}
	s.Show("ages", ages)

	ages = append(ages, 5)
	s.Show("append(ages, 5)", ages)
}
