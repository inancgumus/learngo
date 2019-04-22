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
	"strconv"
	"strings"
)

// domain represents a domain log record
type domain struct {
	name   string
	visits int
}

// parser parses a log file and provides an iterator to iterate upon the domains
//
// the parser struct is carefully crafted to be usable using its zero values except the map field
type parser struct {
	sum     map[string]int // visits per unique domain
	domains []domain       // unique domain names
	total   int            // total visits to all domains
	lines   int            // number of parsed lines (for the error messages)
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	p := parser{
		sum: make(map[string]int),
	}

	// Scan the standard-in line by line
	for in.Scan() {
		p.lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}
		name, visits := fields[0], fields[1]

		// Sum the total visits per domain
		n, _ := strconv.Atoi(visits)
		if n < 0 {
			fmt.Printf("wrong input: %q (line #%d)\n", visits, p.lines)
			return
		}

		d := domain{name: name, visits: n}

		// Collect the unique domains
		if _, ok := p.sum[name]; !ok {
			p.domains = append(p.domains, d)
		}

		p.sum[d.name] += n
		p.total += d.visits
	}

	// Print the visits per domain
	for _, d := range p.domains {
		vis := p.sum[d.name]

		fmt.Printf("%-25s -> %d\n", d.name, vis)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-25s -> %d\n", "TOTAL", p.total)
}
