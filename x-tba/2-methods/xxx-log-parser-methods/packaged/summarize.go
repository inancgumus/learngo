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

	"github.com/inancgumus/learngo/x-tba/2-methods/xxx-log-parser-methods/packaged/metrics"
)

// summarize prints the report and errors if any
func summarize(rep *metrics.Report, errs ...error) {
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	next, cur := rep.Iterator()
	for next() {
		rec := cur()
		fmt.Printf("%-30s %10d\n", rec.Domain, rec.Visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", rep.Total().Visits)

	// only handle the errors once
	dumpErrs(errs...)
}

// this variadic func simplifies the multiple error handling
func dumpErrs(errs ...error) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
		}
	}
}
