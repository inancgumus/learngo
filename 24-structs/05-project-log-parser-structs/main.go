// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// visit stores metrics for a domain
type visit struct {
	domain string
	total  int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]visit // metrics per domain
	domains []string         // unique domain names
	total   int              // total visits for all domains
	lines   int              // number of parsed lines (for the error messages)
}

func main() {
	p := parser{sum: make(map[string]visit)}

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		domain := fields[0]

		// Collect the unique domains
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// Keep track of total and per domain visits
		p.total += visits

		// You cannot assign to composite values
		// p.sum[name].count += visits

		// create and assign a new copy of `visit`
		p.sum[domain] = visit{
			domain: domain,
			total:  visits + p.sum[domain].total,
		}
	}

	// Print the visits per domain
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		visits := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits.total)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}
