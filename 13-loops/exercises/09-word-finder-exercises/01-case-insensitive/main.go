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

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("It needs a keyword input.")
		return
	}
	queries := os.Args[1:]

scan:
	for _, query := range queries {
		query = strings.ToLower(query)
	check:
		for i, word := range strings.Fields(corpus) {
			switch word {
			case "and", "or", "the":
				break check
			}
			if query == word {
				fmt.Printf("#%-2d: %q\n", i+1, word)
				continue scan
			}
		}
	}

}
