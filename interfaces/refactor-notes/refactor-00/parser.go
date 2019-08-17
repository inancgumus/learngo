// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

// parser keeps track of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // the last error occurred

	// totalVisits  int               // total visits for all domains
	// totalUniques int               // total uniques for all domains
}

// newParser constructs, initializes and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parse a log line and return the result
func parse(p *parser, line string) (r result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	r, err := parseResult(line)
	if err != nil {
		p.lerr = fmt.Errorf("line %d: %v", p.lines, err)
	}

	return r
}

// update the parsing results
func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	// Collect the unique domains
	cur, ok := p.sum[r.domain]
	if !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	// p.totalVisits += r.visits
	// p.totalUniques += r.uniques

	// create and assign a new copy of `visit`
	// p.sum[r.domain] = result{
	// 	domain:  r.domain,
	// 	visits:  r.visits + cur.visits,
	// 	uniques: r.uniques + cur.uniques,
	// }
	p.sum[r.domain] = addResult(r, cur)
}

// err returns the last error encountered
func err(p *parser) error {
	return p.lerr
}
