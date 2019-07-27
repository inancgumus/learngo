// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strings"
)

// summarize prints the parsing results.
func summarize(s *summary) {
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for next, cur := iteratorSummary(s); next(); {
		r := cur()
		fmt.Printf("%-30s %10d\n", r.domain, r.visits)
	}

	t := totalsSummary(s)
	fmt.Printf("\n"+"%-30s %10d\n", "TOTAL", t.visits)
}

// this variadic func simplifies the multiple error handling
func dumpErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
		}
	}
}
