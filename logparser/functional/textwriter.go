// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"io"
	"sort"
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
	dash       = "-"
	dashLength = 58
)

func textWriter(w io.Writer) outputFn {
	return func(res []result) error {
		sort.Slice(res, func(i, j int) bool {
			return res[i].domain > res[j].domain
		})

		var total result

		fmt.Fprintf(w, header, "DOMAINS", "PAGES", "VISITS", "UNIQUES")
		fmt.Fprintf(w, strings.Repeat(dash, dashLength)+"\n")

		for _, r := range res {
			total = total.add(r)
			fmt.Fprintf(w, line, r.domain, r.page, r.visits, r.uniques)
		}

		fmt.Fprintf(w, footer, "", total.visits, total.uniques)

		return nil
	}
}

func noWhere() outputFn {
	return func(res []result) error {
		return nil
	}
}
