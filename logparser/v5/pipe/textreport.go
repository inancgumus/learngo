// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

import (
	"fmt"
	"io"
	"text/tabwriter"
)

const (
	minWidth = 0
	tabWidth = 4
	padding  = 4
	flags    = 0
)

// TextReport report generator.
type TextReport struct {
	w io.Writer
}

// NewTextReport returns a TextReport report generator.
func NewTextReport(w io.Writer) *TextReport {
	return &TextReport{w: w}
}

// Consume generates a text report.
func (t *TextReport) Consume(records Iterator) error {
	w := tabwriter.NewWriter(t.w, minWidth, tabWidth, padding, ' ', flags)

	fmt.Fprintf(w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	fmt.Fprintf(w, "-------\t-----\t------\t-------\n")

	var total record

	printLine := func(r Record) error {
		total = r.sum(total)

		fmt.Fprintf(w, "%s\t%s\t%d\t%d\n",
			r.domain, r.page,
			r.visits, r.uniques,
		)

		return nil
	}

	if err := records.Each(printLine); err != nil {
		return err
	}

	fmt.Fprintf(w, "\t\t\t\n")
	fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", "TOTAL", "",
		total.visits,
		total.uniques,
	)

	return w.Flush()
}
