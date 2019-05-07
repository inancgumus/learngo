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

	"github.com/inancgumus/learngo/28-methods/xxx-log-parser-methods/packaged/metrics"
)

// summarize prints the report and errors if any
func summarize(rep *metrics.Report, errs ...error) {
	// TODO: make it strings.Builder

	const format = "%-30s %10s %20s\n"
	const formatValue = "%-30s %10d %20d\n"

	fmt.Printf(format, "DOMAIN", "VISITS", "TIME SPENT")
	fmt.Println(strings.Repeat("-", 65))

	next, cur := rep.Iterator()
	for next() {
		rec := cur()
		fmt.Printf(formatValue, rec.Domain, rec.Visits, rec.TimeSpent)
	}

	fmt.Printf("\n"+formatValue, "TOTAL",
		rep.Total().Visits, rep.Total().TimeSpent,
	)

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
