// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	// declare an uninitialized nil slice
	var uniques []int

loop:
	// you can still use the len function on a nil slice
	for len(uniques) < max {
		n := rand.Intn(max + 1)

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		// let's grow the slice by appending
		uniques = append(uniques, n)
	}

	s.Show("Uniques", uniques)
}
