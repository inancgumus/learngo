// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	ages := []int{35, 15, 25}
	first, last := ages[0:1], ages[1:3]

	s.Show("ages", ages)
	s.Show("first", first)
	s.Show("last", last)

	s.Show("nil slice", []int(nil))
}
