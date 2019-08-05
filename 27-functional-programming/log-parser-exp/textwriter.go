package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: sort by result key interfaces section

func textWriter(w io.Writer, results []result) error {
	fmt.Fprintf(w, "%-25s %-10s %10s %10s\n",
		"DOMAINS", "PAGES", "VISITS", "UNIQUES")

	fmt.Fprintln(w, strings.Repeat("-", 58))

	var total result

	for _, r := range results {
		total = total.add(r)

		fmt.Fprintf(w, "%-25s %-10s %10d %10d\n",
			r.domain, r.page, r.visits, r.uniques)
	}

	fmt.Fprintf(w, "\n%-36s %10d %10d\n",
		"", total.visits, total.uniques)

	return nil
}
