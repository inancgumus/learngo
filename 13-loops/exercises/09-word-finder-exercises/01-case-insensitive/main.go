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

const corpus = "The lazy cat watched the lazy dog sleep all afternoon in the warm, lazy sun."
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



func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please enter the command line arguements:")
		return
	}
	words := strings.Fields(corpus)
	query := os.Args[1:]
queries:
	for _, q := range query {
	// jumpto:
		for i, w := range words {
			b := strings.ToLower(q)
			// switch q {
			// case "and", "or", "it":
			// 	break jumpto
			// }
			if w == b {
				fmt.Printf("#%-2d : %q\n", i+1, b)
				// break queries
				break queries
			}
		}
	}

}
