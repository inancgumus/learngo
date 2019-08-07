package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: sort by result key interfaces section

const (

	//          DOMAINS PAGES VISITS UNIQUES
	//             ^      ^     ^    ^
	//             |      |     |    |
	header     = "%-25s %-10s %10s %10s\n"
	line       = "%-25s %-10s %10d %10d\n"
	footer     = "\n%-36s %10d %10d\n" // -> "" VISITS UNIQUES
	dashLength = 58
)

func textWriter(w io.Writer) outputFunc {
	return func(results []result) error {
		fmt.Fprintf(w, header, "DOMAINS", "PAGES", "VISITS", "UNIQUES")
		fmt.Fprintln(w, strings.Repeat("-", dashLength))

		var total result
		for _, r := range results {
			total = total.add(r)
			fmt.Fprintf(w, line, r.domain, r.page, r.visits, r.uniques)
		}

		fmt.Fprintf(w, footer, "", total.visits, total.uniques)

		return nil
	}
}

func noWhere() outputFunc {
	return func(res []result) error {
		return nil
	}
}
