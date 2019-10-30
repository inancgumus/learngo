// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package report

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/inancgumus/learngo/logparser/v6/logly/record"
)

// Text generates a text report.
func Text(w io.Writer, rs []record.Record) error {
	tw := tabwriter.NewWriter(
		w,
		0,   // minWidth
		4,   // tabWidth
		4,   // padding
		' ', // padChar
		0,   // flags
	)

	fmt.Fprintf(tw, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	fmt.Fprintf(tw, "-------\t-----\t------\t-------\n")

	var total record.Record
	for _, r := range rs {
		total.Sum(r)

		fmt.Fprintf(tw, "%s\t%s\t%d\t%d\n",
			r.Domain, r.Page,
			r.Visits, r.Uniques,
		)
	}

	fmt.Fprintf(tw, "\t\t\t\n")
	fmt.Fprintf(tw, "%s\t%s\t%d\t%d\n",
		"TOTAL", "",
		total.Visits, total.Uniques,
	)

	return tw.Flush()
}
