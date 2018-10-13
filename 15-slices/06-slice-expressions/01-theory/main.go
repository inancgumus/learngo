// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import s "github.com/inancgumus/prettyslice"

func main() {
	slicable := []byte{'c', 'o', 'l', 'o', 'r'}
	s.Show("slicable[0:5]", slicable[0:5])
	s.Show("slicable[0:2]", slicable[0:2])
	s.Show("slicable[3:5]", slicable[3:5])
	s.Show("slicable[1:4]", slicable[1:4])
}
