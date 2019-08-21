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
	"strconv"
	"strings"
)

// result stores details about a log line
type result struct {
	domain string
	visits int
	// add more metrics when needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // the last error occurred
}

// newParser creates and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parse result from a log line
func (p *parser) parse(line string) (r result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error

	r.domain = fields[0]
	r.visits, err = strconv.Atoi(fields[1])

	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}
	return
}

// update the parsing results
func (p *parser) update(r result) {
	if p.lerr != nil {
		return
	}

	// collect unique domains for easy access to sum map
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// keep track of total visits
	p.total += r.visits

	// group the log lines by domain
	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}

// summarize the parsing results
func (p *parser) summarize() {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

// err returns the last error encountered
func (p *parser) err() error {
	return p.lerr
}
