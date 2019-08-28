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

	"github.com/inancgumus/learngo/logparser/v6/pipe"
)

const (
	minWidth = 0
	tabWidth = 4
	padding  = 4
	flags    = 0
)

// Text report generator.
type Text struct {
	w io.Writer
}

// AsText returns a Text report generator.
func AsText(w io.Writer) *Text {
	return &Text{w: w}
}

// Consume generates a text report.
func (t *Text) Consume(records pipe.Iterator) error {
	w := tabwriter.NewWriter(t.w, minWidth, tabWidth, padding, ' ', flags)

	write := fmt.Fprintf

	write(w, "DOMAINS\tPAGES\tVISITS\tUNIQUES\n")
	write(w, "-------\t-----\t------\t-------\n")

	var total pipe.Record

	records.Each(func(r pipe.Record) {
		if r, ok := r.(pipe.Summer); ok {
			total = r.Sum(total)
		}

		write(w, "%s\t%s\t%d\t%d\n",
			r.Str("domain"), r.Str("page"),
			r.Int("visits"), r.Int("uniques"),
		)
	})

	write(w, "\t\t\t\n")
	write(w, "%s\t%s\t%d\t%d\n", "TOTAL", "",
		total.Int("visits"),
		total.Int("uniques"),
	)

	return w.Flush()
}
