package main

import (
	"fmt"
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
	err     error          // saves the last error occurred
}

// newParser creates and returns a new parser.
func newParser() *parser {
	return &parser{sum: make(map[string]int)}
}

func add(p *parser, line string) {
	// if there was a previous error do not add
	if p.err != nil {
		return
	}

	dom, err := parse(p, line)

	// store only the last error
	if err != nil {
		p.err = err
		return
	}

	push(p, dom)
}

// iterator returns two functions for iterating over domains.
// next = returns true when there are more domains to iterate on.
// cur  = returns the current domain
//
// READ METHOD
func iterator(p *parser) (next func() bool, cur func() domain) {
	// remember the last received line
	var last int

	next = func() bool {
		defer func() { last++ }()
		return len(p.domains) > last
	}

	cur = func() domain {
		d := p.domains[last-1]
		vis := p.sum[d.name]

		// return a copy so the caller cannot change it
		return domain{name: d.name, visits: vis}
	}

	return
}

// error returns the last error occurred
//
// READ METHOD
func err(p *parser) error {
	return p.err
}

func parse(p *parser, line string) (dom domain, err error) {
	p.lines++ // increase the parsed line counter (only write is here)

	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return dom, err
	}

	name, visits := fields[0], fields[1]

	n, err := strconv.Atoi(visits)
	if n < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", visits, p.lines)
		return dom, err
	}

	return domain{name: name, visits: n}, nil
}

func push(p *parser, d domain) {
	// collect the unique domains
	if _, ok := p.sum[d.name]; !ok {
		p.domains = append(p.domains, d)
	}

	p.sum[d.name] += d.visits
	p.total += d.visits
}
