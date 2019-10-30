// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/inancgumus/learngo/logparser/testing/report"
)

// summarize prints the parsing results.
//
// it prints the errors and returns if there are any.
//
// --json flag encodes to json and prints.
func summarize(sum *report.Summary, errors ...error) {
	if errs(errors...) {
		return
	}

	if args := os.Args[1:]; len(args) == 1 && args[0] == "--json" {
		encode(sum)
		return
	}
	stdout(sum)
}

// encodes the summary to json
func encode(sum *report.Summary) {
	out, err := json.MarshalIndent(sum, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(out)
}

// prints the summary to standard out
func stdout(sum *report.Summary) {
	const (
		head = "%-30s %10s %20s\n"
		val  = "%-30s %10d %20d\n"
	)

	fmt.Printf(head, "DOMAIN", "VISITS", "TIME SPENT")
	fmt.Println(strings.Repeat("-", 65))

	for next, cur := sum.Iterator(); next(); {
		r := cur()
		fmt.Printf(val, r.Domain, r.Visits, r.TimeSpent)
	}

	t := sum.Total()
	fmt.Printf("\n"+val, "TOTAL", t.Visits, t.TimeSpent)
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
