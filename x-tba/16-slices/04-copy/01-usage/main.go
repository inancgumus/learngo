// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	me := []byte{'m', 'e'}
	yo := []byte{'y', 'o', '!'}

	s.Show("me [before]", me)
	s.Show("yo [before]", yo)

	N := copy(me, yo)
	fmt.Printf("%d element(s) copied.\n", N)

	s.Show("me [after]", me)
	s.Show("yo [after]", yo)
}
