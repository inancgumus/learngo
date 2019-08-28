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

type textReport struct{}

func newTextReport() *textReport {
	return new(textReport)
}

func (s *textReport) digest(records iterator) error {
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, ' ', flags)

	write := fmt.Fprintf

	write(w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	write(w, "-------\t-----\t------\t-------\n")

	var total record
	records.each(func(r record) {
		total = total.sum(r)

		write(w, "%s\t%s\t%d\t%d\n", r.domain, r.page, r.visits, r.uniques)
	})

	write(w, "\t\t\t\n")
	write(w, "%s\t%s\t%d\t%d\n", "TOTAL", "", total.visits, total.uniques)

	return w.Flush()
}
