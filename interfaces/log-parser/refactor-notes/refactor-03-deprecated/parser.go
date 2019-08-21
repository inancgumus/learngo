// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// parser keeps track of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
}

// newParser constructs, initializes and returns a new parser
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parse all the log lines and update the results
func parse(p *parser) error {
	process := func(r result) {
		update(p, r)
	}

	err := scan(process)

	return err
}

func update(p *parser, r result) {
	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// create and assign a new copy of `visit`
	p.sum[r.domain] = addResult(r, p.sum[r.domain])
}
