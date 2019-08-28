// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

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

	write := fmt.Fprintf

	write(w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	write(w, "-------\t-----\t------\t-------\n")

	var total Record

	err := records.Each(func(r Record) error {
		total = r.Sum(total)

		write(w, "%s\t%s\t%d\t%d\n",
			r.Domain, r.Page,
			r.Visits, r.Uniques,
		)

		return nil
	})
	if err != nil {
		return err
	}

	write(w, "\t\t\t\n")
	write(w, "%s\t%s\t%d\t%d\n", "TOTAL", "",
		total.Visits,
		total.Uniques,
	)

	return w.Flush()
}
