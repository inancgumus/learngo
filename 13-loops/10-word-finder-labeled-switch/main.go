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
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	// labels and other names do not share the same scope
	// this will work even though `queries` label exists
	//
	// var queries string
	// _ = queries

	// this label labels the parent loop below.
	// label's scope is the whole func main()
queries:
	for _, q := range query {

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				continue queries
			}
		}
	}
}
