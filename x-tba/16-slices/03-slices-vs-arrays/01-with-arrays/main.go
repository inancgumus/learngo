// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 10
	// max, _ := strconv.Atoi(os.Args[1])

	var uniques [max]int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max + 1)

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques[found] = n
		found++
	}

	s.Show("Uniques", uniques)
}
