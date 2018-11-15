// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = `lazy cat jumps again and again and again...
since she is very excited and happy.`

func main() {
	query := os.Args[1:]
	words := strings.Fields(corpus)

	filter := [...]string{
		"and", "or", "the", "is", "since", "very",
	}

queries:
	for _, q := range query {
		for _, v := range filter {
			if q == v {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}
