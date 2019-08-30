// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package report

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/inancgumus/learngo/logparser/v6/parse"
)

// TextReport report generator.
type TextReport struct {
	w     *tabwriter.Writer
	total parse.Record
}

// Text creates a report generator.
func Text(w io.Writer) *TextReport {
	tw := tabwriter.NewWriter(w,
		0,   // minWidth
		4,   // tabWidth
		4,   // padding
		' ', // padChar
		0,   // flags
	)

	return &TextReport{w: tw}
}

// Generate the report from the records.
// tabwriter caches, the memory usage will be high
// if you send a large number of records to Generate.
func (tr *TextReport) Generate(rs []parse.Record) error {
	fmt.Fprintf(tr.w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	fmt.Fprintf(tr.w, "-------\t-----\t------\t-------\n")

	for _, r := range rs {
		tr.total.Sum(r)

		fmt.Fprintf(tr.w, "%s\t%s\t%d\t%d\n",
			r.Domain, r.Page,
			r.Visits, r.Uniques,
		)
	}

	fmt.Fprintf(tr.w, "\t\t\t\n")
	fmt.Fprintf(tr.w, "%s\t%s\t%d\t%d\n",
		"TOTAL", "",
		tr.total.Visits, tr.total.Uniques,
	)

	return tr.w.Flush()
}
