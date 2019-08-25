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
)

// parser keeps track of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
}

// newParser constructs, initializes and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parse the log lines and return results
func parse(p *parser) error {
	var (
		l  = 1
		in = bufio.NewScanner(os.Stdin)
	)

	for in.Scan() {
		r, err := parseResult(in.Text())
		if err != nil {
			return fmt.Errorf("line %d: %v", l, err)
		}

		l++

		update(p, r)
	}

	return in.Err()
}

// update the parsing results
func update(p *parser, r result) {
	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// create and assign a new copy of `visit`
	p.sum[r.domain] = addResult(r, p.sum[r.domain])
}
