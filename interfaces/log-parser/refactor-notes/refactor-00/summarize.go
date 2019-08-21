// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"sort"
	"strings"
)

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

// summarize summarizes and prints the parsing result
func summarize(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf(header, "DOMAIN", "PAGES", "VISITS", "UNIQUES")
	fmt.Println(strings.Repeat("-", dashLength))

	var total result

	for _, domain := range p.domains {
		r := p.sum[domain]
		total = addResult(total, r)

		fmt.Printf(line, r.domain, r.page, r.visits, r.uniques)
	}
	fmt.Printf(footer, "TOTAL", total.visits, total.uniques)
}
