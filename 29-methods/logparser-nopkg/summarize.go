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
//
// it prints the errors and returns if there are any.
//
// --json flag encodes to json and prints.
func summarize(sum *summary, errors ...error) {
	if errs(errors...) {
		return
	}

	const (
		head = "%-30s %10s %20s\n"
		val  = "%-30s %10d %20d\n"
	)

	fmt.Printf(head, "DOMAIN", "VISITS", "TIME SPENT")
	fmt.Println(strings.Repeat("-", 65))

	for next, cur := sum.iterator(); next(); {
		r := cur()
		fmt.Printf(val, r.domain, r.visits, r.timeSpent)
	}

	t := sum.totals()
	fmt.Printf("\n"+val, "TOTAL", t.visits, t.timeSpent)
}

// this variadic func simplifies the multiple error handling
func errs(errs ...error) (wasErr bool) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
			wasErr = true
		}
	}
	return
}
