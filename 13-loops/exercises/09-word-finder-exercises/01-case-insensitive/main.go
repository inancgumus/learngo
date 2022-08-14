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

const corpus = "lazy cat jumps again and again and again lazy"

func main() {
	words := strings.Fields(corpus)
	args := os.Args[1:]
	var query string
	if len(args) != 0 {
		query = strings.ToLower(args[0])
	} else {
		fmt.Println("Not enough arguments")
	}

	for p, w := range words {
		if w == query {
			fmt.Println(p, w)
		}

	}
}
