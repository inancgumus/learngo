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
	"text/tabwriter"
)

// TODO: make this configurable? or exercise?
const (
	minWidth = 0
	tabWidth = 4
	padding  = 4
	flags    = 0
)

type textSummary struct{}

func newTextSummary() *textSummary {
	return new(textSummary)
}

func (s *textSummary) summarize(results iterator) error {
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, ' ', flags)

	write := fmt.Fprintf

	write(w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	write(w, "-------\t-----\t------\t-------\n")

	var total result
	results.each(func(r result) {
		total = total.add(r)

		write(w, "%s\t%s\t%d\t%d\n", r.domain, r.page, r.visits, r.uniques)
	})

	write(w, "\t\t\t\n")
	write(w, "%s\t%s\t%d\t%d\n", "TOTAL", "", total.visits, total.uniques)

	return w.Flush()
}
