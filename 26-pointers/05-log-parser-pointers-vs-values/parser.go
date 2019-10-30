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
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // the last error occurred
}

// newParser constructs, initializes and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parse parses a log line and returns the parsed result with an error
func parse(p *parser, line string) (r result) {
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

// update updates all the parsing results using the given parsing result
func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	p.total += r.visits

	// create and assign a new copy of `visit`
	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
}

// err returns the last error encountered
func err(p *parser) error {
	return p.lerr
}
